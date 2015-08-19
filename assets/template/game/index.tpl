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
    padding : 0px;
    margin : 0px;
    width  : 30px;
    height : 30px;
}
div.grid {
    padding : 0px;
    margin : 0px;
    width  : 30px;
    height : 30px;
}

div.me {
    background-color: lightgreen;
}

div.enemy {
    background-color: lightblue;
}

</style>

潜水艦ゲーム

<div id="status-container"></div>
<div id="game-container">
</div>

<script type="text/html" id="tmpl_game">

<table><tr><td>
    <center>自軍</center>
    <div id="my-map" style="margin:10px">
        <table class="grid"> 
        <% for (y=0;y<Me["Fields"].length;++y) { %>
            <tr>
            <% for (x=0;x<Me["Fields"][y].length;++x) { %>
                <td>
                <div id="me_<%= y %>_<%= x %>" class="grid me" data-hit-type="<%= Me["Fields"][y][x]["HitType"] %>">
                <% if ( Me["Fields"][y][x]["ShipID"] != 0 ) { %>
                    <img width="30px" height="30px" src="/static/images/ship_<%= Me["Fields"][y][x]["ShipPart"] %>.png"
                    <% if( Me["Fields"][y][x]["ShipDirection"] ) { %> 
                        style="transform: rotate(-90deg);" 
                     <% } %> 
                    />
                <% } else { %>
                <% }  %>
                </div></td>
           <% } %>
           </tr>
        <% } %>
        </table> 
    </div>
    </td><td>
    <center>敵軍</center>
    <div id="enemy-map" style="margin:10px">
        <table class="grid"> 
        <% for (y=0;y<Enemy["Fields"].length;++y) { %>
            <tr>
            <% for (x=0;x<Enemy["Fields"][y].length;++x) { %>
                <td><div id="enemy_<%= y %>_<%= x %>" class="grid enemy" data-hit-type="<%= Enemy["Fields"][y][x]["HitType"] %>" onClick="game.attack(<%= y %>,<%= x %>);"></div></td>
           <% } %>
           </tr>
        <% } %>
        </table> 
    </div>
</td></tr></table>


</script>
<script type="text/html" id="tmpl_game2">
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
    user_id : null,
    matching_id : null,
    is_your_turn : false,
    start : function(matching_id,user_id){
        game.matching_id =  matching_id;
        game.user_id =  user_id;
        game._connect();
        game._start();
    },
    attack : function(y,x) {
        // 自分のターンの場合のみ
            
        if ( game.is_your_turn ) {
            var hit_type = $('#enemy_' + y + '_' + x).attr("data-hit-type");

            // 攻撃をまだしていないフィールド
            if ( hit_type  == 0  ){
                
                game.socket.send('{"cmd":"attack","matching_id":"' + game.matching_id + '","y" : ' + y + ',"x" : ' + x + '}');

            }
          
        }
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
            if (event && event.data) {
                var data =JSON.parse(event.data)

                if (data["cmd"] == "start" ) {
                    game.is_your_turn =data["data"]["IsYourTurn"];
                    $('#game-container').html( $('#tmpl_game').template(data["data"]) );
                    if ( game.is_your_turn ) {
                        $('#status-container').html("あなたのターンです");
                    }
                    else {
                        $('#status-container').html("敵ののターンです");
                    }
                        
                }
                else if (data["cmd"] == "turn_end" ) {
                    game.is_your_turn = false;
                    id = '#enemy_' + data["data"]["y"] + '_' + data["data"]["x"];
                    field = data["data"]["field"];
                    $(id).attr("data-hit-type",field["HitType"]);
                    
                    if ( field["ShipID"] != 0 ) {
                        html = '<img width="30px" height="30px" src="/static/images/ship_' +  
                            field["ShipPart"] + '.png"';
                        if ( field["ShipDirection"] ) {
                            html += ' style="transform: rotate(-90deg);"'
                        }
                        html += ">";

                        $(id).html(html);
                    }

                    // 近く
                    if ( field["HitType"] == 2 ) {
                        $(id).css("background-color","pink");
                    }
                    else {
                        $(id).css("background-color","red");
                    }

                    if ( data["data"]["on_finish"] ) {
                        game.is_your_turn = false;
                        $('#status-container').html("勝利！");
                    }
                    else {
                        $('#status-container').html("敵ののターンです");
                    }
                }
                else if (data["cmd"] == "turn_start" ) {
                    game.is_your_turn = true;
                    id = '#me_' + data["data"]["y"] + '_' + data["data"]["x"];
                    $(id).attr("data-hit-type",data["data"]["field"]["HitType"]);
                    $(id).css("background-color","red");

                    if ( data["data"]["on_finish"] ) {
                        game.is_your_turn = false;
                        $('#status-container').html("敗北！");
                    }
                    else {
                        $('#status-container').html("あなたのターンです");
                    }

                }
            }
        };
    },
    _start : function(){
        // need to way to open
        game.socket.onopen = function() { 
            game.socket.send('{"cmd":"start","matching_id":"' + game.matching_id + '","user_id":"' + game.user_id +'"}');
        };


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
            if (event && event.data) {
                var data =JSON.parse(event.data)

                if (data["cmd"] == "found" ) {
                    $('#status-container').html("Found!");
                    game.start( data["matching_id"] , data["user_id"]);
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
