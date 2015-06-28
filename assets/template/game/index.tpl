<html>
<head>
<title>submarine game</title>
</head>
<body>

潜水艦ゲーム


<script type="text/javascript">
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
            if (event && event.data) {
                console.log(event.data);
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
