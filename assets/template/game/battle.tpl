<html>
<head>
<title>submarine game</title>
<link rel="stylesheet" type="text/css" href="/static/css/base.css" />
</head>
<body>

<style>
table.ships {
}
table.ships th {
    background-color : lightyellow;
}
table.ships td, th {
    padding : 0px;
    margin : 0px;
    width  : 10px;
    height : 10px;
}

table.grid td, th {
    border: 2px #ffffff solid;
    padding : 0px;
    margin : 0px;
    width  : 25px;
    height : 25px;
}
div.grid {
    padding : 0px;
    margin : 0px;
    width  : 25px;
    height : 25px;
    background-image: url("/static/images/sea.png");
    background-size:cover;
}

</style>

潜水艦ゲーム

<div id="status-container"></div>
<div id="game-container">
</div>

<script type="text/html" id="tmpl_game">

<table><tr><td>
    <div id="my-map" style="margin:10px">
    <center>自軍</center>

    <table class="ships">
    <tr>
    <th>&nbsp;</th>
    <th style="width:100">残数</th>
    </tr>
    <% for ( i = 0 ; i < Ships.length ; i++ ) { %>
    <tr>
        <td>
        <table><tr>
        <% for ( j = 0 ; j < Ships[i]["size"] ; j++ ) { %>
            <% part = 0 ; %>
            <% if ( j == Ships[i]["size"] - 1 ) { part = 2 } else if (j == 0 ) { part = 0 } else { part = 1 } %>
            <td><img width="10px" height="10px" src="/static/images/ship_<%= part %>.png" 
                style="-moz-transform:rotate(-90deg); -webkit-transform:rotate(-90deg); transform: rotate(-90deg);" /></td>
        <% } %>
        </tr>
        </table>
        </td>
        <td align="right"><span id="me_ship_<%=Ships[i]["size"]%>"><%= Ships[i]["count"] %></span></td>
    </tr>
    <% } %>
    </table>


        <table class="grid"> 
        <% for (y=0;y<Me["Fields"].length;++y) { %>
            <tr>
            <% for (x=0;x<Me["Fields"][y].length;++x) { %>
                <td>
                <div id="me_<%= y %>_<%= x %>" class="grid me" data-hit-type="<%= Me["Fields"][y][x]["HitType"] %>">
                <% if ( Me["Fields"][y][x]["ShipID"] != 0 ) { %>
                    <img width="25px" height="25px" src="/static/images/ship_<%= Me["Fields"][y][x]["ShipPart"] %>.png"
                    <% if( Me["Fields"][y][x]["ShipDirection"] ) { %> 
                        style="-moz-transform:rotate(-90deg); -webkit-transform:rotate(-90deg);transform: rotate(-90deg);" 
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

    <div id="enemy-map" style="margin:10px">


    <center>敵軍</center>


    <table class="ships">
    <tr>
    <th>&nbsp;</th>
    <th style="width:100">残数</th>
    </tr>
    <% for ( i = 0 ; i < Ships.length ; i++ ) { %>
    <tr>
        <td>
        <table><tr>
        <% for ( j = 0 ; j < Ships[i]["size"] ; j++ ) { %>
            <% part = 0 ; %>
            <% if ( j == Ships[i]["size"] - 1 ) { part = 2 } else if (j == 0 ) { part = 0 } else { part = 1 } %>
            <td><img width="10px" height="10px" src="/static/images/ship_<%= part %>.png" style="-moz-transform:rotate(-90deg); -webkit-transform:rotate(-90deg); transform: rotate(-90deg);" /></td>
        <% } %>
        </tr>
        </table>
        </td>
        <td align="right"><span id="enemy_ship_<%=Ships[i]["size"]%>"><%= Ships[i]["count"] %></span></td>
    </tr>
    <% } %>
    </table>


        <table class="grid" id="enemy-grid"> 
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

<script type="text/html" id="tmpl_grid">
    <% for (y=0;y<enemyFields.length;++y) { %>
    <tr>
	<% for (x=0;x<enemyFields[y].length;++x) { %>
	<td>
	    <div id="enemy_<%= y %>_<%= x %>" class="grid enemy" data-hit-type="<%= enemyFields[y][x]["HitType"] %>" onClick="game.attack(<%= y %>,<%= x %>);">
	        <% if (enemyFields[y][x]["HitType"] == 0 && enemyFields[y][x]["ShipID"] != 0) { %>
		    <img width="25px" height="25px" src="/static/images/ship_<%= enemyFields[y][x]["ShipPart"] %>.png"
		        <% if( enemyFields[y][x]["ShipDirection"] ) { %>
			    style="-moz-transform:rotate(-90deg); -webkit-transform:rotate(-90deg);transform: rotate(-90deg);"
		        <% } %>
		    />
		<% } else if (enemyFields[y][x]["HitType"] == 1) { %>
		    <img width="25px" height="25px" src="/static/images/ship_broken_<%= enemyFields[y][x]["ShipPart"] %>.png"
		        <% if( enemyFields[y][x]["ShipDirection"] ) { %>
		            style="-moz-transform:rotate(-90deg); -webkit-transform:rotate(-90deg);transform: rotate(-90deg);"
		        <% } %>
		    />
		<% } else if (enemyFields[y][x]["HitType"] == 2) { %>
		    <img width="25px" height="25px" src="/static/images/bomb_near.png">
		<% } else if (enemyFields[y][x]["HitType"] == 3) { %>
		    <img width="25px" height="25px" src="/static/images/bomb_miss.png">
		<% } %>
	    </div>
	</td>
	<% } %>
    </tr>
    <% } %>
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
                        $('#enemy-map').css('background-color','yellow');
                    }
                    else {
                        $('#status-container').html("敵ののターンです");
                    }
                        
                }
                else if (data["cmd"] == "turn_end" ) {

                    $('#enemy-map').css('background-color','#ffffff');



                    game.is_your_turn = false;
                    id = '#enemy_' + data["data"]["y"] + '_' + data["data"]["x"];
                    field = data["data"]["field"];
                    $(id).attr("data-hit-type",field["HitType"]);
                    
                    if ( field["ShipID"] != 0 ) {
                        html = '<img width="25px" height="25px" src="/static/images/ship_broken_' +  
                            field["ShipPart"] + '.png"';
                        if ( field["ShipDirection"] ) {
                            html += ' style="-moz-transform:rotate(-90deg); -webkit-transform:rotate(-90deg); transform: rotate(-90deg);"'
                        }
                        html += ">";

                        $(id).html(html);

                    } else if ( field["HitType"] == 2 ) {
                        html = '<img width="25px" height="25px" src="/static/images/bomb_near.png">';
                        $(id).html(html);
                    }
                    else {
                        html = '<img width="25px" height="25px" src="/static/images/bomb_miss.png">';
                        $(id).html(html);
                    }

                    // 船撃沈
                    if ( data["data"]["destroy_ship_size"]  != 0 ) {
                        num_id = '#enemy_ship_' + data["data"]["destroy_ship_size"]  ;
                        count = $(num_id).text();
                        $(num_id).text( count - 1);
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

                    $('#enemy-map').css('background-color','yellow');

                    game.is_your_turn = true;
                    id = '#me_' + data["data"]["y"] + '_' + data["data"]["x"];

                    //$(id).attr("data-hit-type",data["data"]["field"]["HitType"]);
                    //$(id).css("background-color","red");

                    var field = data["data"]["field"];

                    if ( field["ShipID"] != 0 ) {
                        html = '<img width="25px" height="25px" src="/static/images/ship_broken_' +  
                            field["ShipPart"] + '.png"';
                        if ( field["ShipDirection"] ) {
                            html += ' style="-moz-transform:rotate(-90deg); -webkit-transform:rotate(-90deg); transform: rotate(-90deg);"'
                        }
                        html += ">";

                        $(id).html(html);
                    } else if ( field["HitType"] == 2 ) {
                        html = '<img width="25px" height="25px" src="/static/images/bomb_near.png">';
                        $(id).html(html);
                    }
                    else {
                        html = '<img width="25px" height="25px" src="/static/images/bomb_miss.png">';
                        $(id).html(html);
                    }



                    // 船撃沈
                    if ( data["data"]["destroy_ship_size"]  != 0 ) {
                        num_id = '#me_ship_' + data["data"]["destroy_ship_size"]  ;
                        count = $(num_id).text();
                        $(num_id).text( count - 1);
                    }

                    if ( data["data"]["on_finish"] ) {
                        game.is_your_turn = false;
                        $('#enemy-grid').html( $('#tmpl_grid').template(data["data"]) );
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
                    
$(document).ready(function(){
    var h = location.hash;
    h = h.replace( /^#/ , '' );
    var tmp = h.split("/");
    console.log(tmp);
    game.start( tmp[0] , tmp[1]);
});

</script>


</body>
</html>