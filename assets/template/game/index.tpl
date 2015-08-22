<html>
<head>
<title>submarine game</title>
<link rel="stylesheet" type="text/css" href="/static/css/base.css" />
</head>
<body>


潜水艦ゲーム

<div id="status-container"></div>

<script src="/static/js/jquery-1.11.2.min.js"></script>
<script src="/static/js/jquery.ze.js"></script>
<script type="text/javascript">

var matching = { 
    socket : null,
    start : function(){

        $('#status-container').html("Searching Enemy...");

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
                var data =JSON.parse(event.data)

                if (data["cmd"] == "found" ) {
                    $('#status-container').html("Found!");
                    location.href = "/game/battle#" +  data["matching_id"] + "/" + data["user_id"];
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
