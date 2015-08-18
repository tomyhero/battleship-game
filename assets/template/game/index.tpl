<html>
<head>
<title>submarine game</title>
</head>
<body>

<style>
html, body, div, span, applet, object, iframe,
    h1, h2, h3, h4, h5, h6, p, blockquote, pre,
    a, abbr, acronym, address, big, cite, code,
    del, dfn, em, img, ins, kbd, q, s, samp,
    small, strike, strong, sub, sup, tt, var,
    b, u, i, center,
    dl, dt, dd, ol, ul, li,
    fieldset, form, label, legend,
    table, caption, tbody, tfoot, thead, tr, th, td,
    article, aside, canvas, details, embed, 
    figure, figcaption, footer, header, hgroup, 
    menu, nav, output, ruby, section, summary,
    time, mark, audio, video {
margin: 0;
padding: 0;
border: 0;
        font-size: 100%;
font: inherit;
      vertical-align: baseline;
    }
/* HTML5 display-role reset for older browsers */
article, aside, details, figcaption, figure, 
    footer, header, hgroup, menu, nav, section {
display: block;
    }
body {
    line-height: 1;
}
ol, ul {
    list-style: none;
}
blockquote, q {
quotes: none;
}
blockquote:before, blockquote:after,
    q:before, q:after {
content: '';
content: none;
    }
table {
    border-collapse: collapse;
    border-spacing: 0;
}

table.grid td, th {
    border: 2px #ffffff solid;
}
div.grid {
    width  : 30px;
    height : 30px;
}

div.me {
    background-color: green;
}

div.enemy {
    background-color: red;
}

</style>

潜水艦ゲーム

<div id="status-container"></div>
<div id="game-container" style="align:top">
</div>

<script type="text/html" id="tmpl_game">
<table><tr><td>
    <center>自軍</center>
    <div id="my-map" style="margin:10px">
        <table class="grid"> 
        <% for (y=0;y<grid.y;++y) { %>
            <tr>
            <% for (x=0;x<grid.x;++x) { %>
                <td><div id="me_<%= y %>_<%= x %>" class="grid me"></div></td>
           <% } %>
           </tr>
        <% } %>
        </table> 
    </div>
    </td><td>
    <center>敵軍</center>
    <div id="enemy-map" style="margin:10px">
        <table class="grid"> 
        <% for (y=0;y<grid.y;++y) { %>
            <tr>
            <% for (x=0;x<grid.x;++x) { %>
                <td><div id="enemy_<%= y %>_<%= x %>" class="grid enemy"></div></td>
           <% } %>
           </tr>
        <% } %>
        </table> 
    </div>
</td></tr></table>
</script>

<script src="/static/js/jquery-1.11.2.min.js"></script>
<script src="/static/js/jquery.ze.js"></script>
<script type="text/javascript">


var game = { 
    socket : null,
    grid : { x : 16 , y : 16 },
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

        $('#game-container').html( $('#tmpl_game').template({ grid : game.grid }) );

        $('#status-container').html("Start Game");
    },
    createGrid : function(prefix){
        var html = "<table>";
        var y,x;
        for (y=0;y<game.grid.y;++y) {
            html += "<tr>";
            for (x=0;x<game.grid.x;++x) {
                html += '<td><div id="' + prefix + '_' + y + '_' + x + '" class="grid ' + prefix + '">#</div></td>';
           }
            html += "</tr>";
        }
        html += "</table>";
        return html;
    }
};

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
        console.log("Called onMessage");
            if (event && event.data) {
                var data =JSON.parse(event.data)
                console.log(data);

                if (data["cmd"] == "found" ) {
                    $('#status-container').html("Found!");
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
