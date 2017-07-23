$(function() {
var sock = (function(form) {
    var socket = null;
        var count = 0;
        var max_event = 9,
        $chatroom = $("#chatroom");

    function connect() {
        if ( isConnected() ) {
            addMessage({
                sender: "server",
                message: "Connect is already open"
            });
        } else {
            var url = "ws://" + document.location.host + "/ws?username=" + $("input[name=username]", $(form) ).val();
            socket = new WebSocket(url );
            socket.onclose = function( e ) {
                $( ".on-chat" ).prop( "disabled", true );
                $( ".on-chat-reverse" ).prop( "disabled", false );
                $( ".btn-connect" ).prop( "disabled", false );
                $( ".btn-disconnect" ).prop( "disabled", true );
                addMessage({
                    sender: "server",
                    message: "Disconnected"
                });
            };
            socket.onopen = function (e) {
                $( ".on-chat" ).prop("disabled", false );
                $( ".on-chat-reverse" ).prop("disabled", true );
                $( ".btn-connect" ).prop( "disabled", true );
                $( ".btn-disconnect" ).prop( "disabled", false );
                $( "input[name=message]", $(form) ).focus();
            };
            socket.onmessage = function( e ) {
                var c = Cookies.get("wsid");
                console.log(Cookies.get());
                console.log(e.data);
                var event = JSON.parse( e.data );
                addMessage( event );
            };
        }
    }
    function adjustScroll() {
        $chatroom.scrollTop( $chatroom.prop( "scrollHeight" ) );
    }
    function addMessage( event ) {
        var $namecard = getNamecard( event.sender );
        var $p = $( "<p/>" );
        $p.append( getNamecard( event.sender ) ).append( event.message );
        $(".chatroom").append( $p );
        adjustScroll();
        count++;
        if ( count > max_event ) {
            $("p", $chatroom).first().remove();
        }
    }
    function getNamecard( name ) {
        //return "[" + name + "] ";
        return $("<button/>", {
            type: "button",
            class: "btn btn-default btn-xs mr5"
        }).text( name );
    }
    function disconnect() {
        if ( socket !== null ) {
            socket.close();
        }
    }
    function isConnected() {
        if ( socket !== null && socket.readyState === socket.OPEN ) {
            return true;
        } else {
            return false;
        }
    }
    function send( message ) {
        var event = {
            message: message
        }
        socket.send( JSON.stringify( event ) );
    }
    return {
        connect: function() {
            connect();
            return isConnected();
        },
        disconnect: function() {
            disconnect();
            return isConnected();
        },
        isConnected: function() {
            return isConnected();
        },
        send: function(msg) {
            if (msg === undefined) {
                send( $("input[name=message]", $(form) ).val() );
                $( "input[name=message]", $(form) ).val("");
            } else {
                send( msg );
            }
        },
        getCount: function() {
            return count;
        }
    }
});


$( ".on-chat" ).prop( "disabled", true );
//$( ".on-chat" ).prop( "disabled", true );

var ws = sock( "#form-chat" );
console.log(ws.isConnected());
$( ".btn-connect" ).click(function( e ) {
    ws.connect();
});
$( ".btn-disconnect" ).click(function( e ) {
    ws.disconnect();
});
$( ".btn-send" ).click(function( e ) {
    e.preventDefault();
    ws.send();
});
$( ".btn-count" ).click(function( e ) {
    e.preventDefault();
    $(this).text("Count: " + ws.getCount())
});
$( ".btn-send-test" ).click(function( e ) {

    for (var i=0; i<100; i++) {
        ws.send(i + " = " + Math.random().toString(36).substring(7));
//        ws.send("test");
    }
});

});
