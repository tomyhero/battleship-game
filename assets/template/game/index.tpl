<html>
<head>
<title>submarine game</title>
<script type="text/javascript" src="/static/js/enchant.js"></script>
</head>
<body>

潜水艦ゲーム


<script type="text/javascript">


enchant();


var game = { 
    socket : null,
    start : function(){
        game._connect();
        game._start();
    },
    _connect : function(){
        var url = "{{.game_endpoint}}";

        // FireFoxとの互換性を考慮してインスタンス化
            
        if ("MozWebSocket" in window) {
            game.socket = new MozWebSocket(url);
        }
        else {
            game.socket = new WebSocket(url);
        }
        console.log("connected");

        game.socket.onmessage = function(event) {
        console.log("Called onMessage");
            if (event && event.data) {
                var data =JSON.parse(event.data)
                console.log(data);
            }
        };
    },
    _start : function(){
        // need to way to open
        game.socket.onopen = function() { 
            game.socket.send('{"cmd":"start"}');
        };
    }

};
var matching = { 
    socket : null,
    start : function(){
        matching._connect();
        matching._search();
    },
    _connect : function(){
        var url = "{{.matching_endpoint}}";

        // FireFoxとの互換性を考慮してインスタンス化
            
        if ("MozWebSocket" in window) {
            matching.socket = new MozWebSocket(url);
        }
        else {
            matching.socket = new WebSocket(url);
        }

        matching.socket.onmessage = function(event) {
        console.log("Called onMessage");
            if (event && event.data) {
                var data =JSON.parse(event.data)
                console.log(data);

                if (data["cmd"] == "found" ) {
                    game.start();
                }

            }
        };
    },
    _search : function(){
        // need to way to open
        matching.socket.onopen = function() { 
            matching.socket.send('{"cmd":"search"}');
        };
    }
};

matching.start();

</script>


</body>
</html>
