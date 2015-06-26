<html>
<head>
<title>submarine game</title>
</head>
<body>

潜水艦ゲーム


<script type="text/javascript">
var game = { 
    connect : function(){
        var url = "{{.matching_endpoint}}";
        var ws;

        // FireFoxとの互換性を考慮してインスタンス化
        if ("WebSocket" in window) {
            ws = new WebSocket(url);
        } else if ("MozWebSocket" in window) {
            ws = new MozWebSocket(url);
        }

        return ws;
    }
};

 var ws = game.connect();

</script>

</body>
</html>
