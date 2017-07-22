$(function() {

var sock = (function(form) {
    var socket = null;

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
                $(".on-chat").prop("disabled", true);
                $(".on-chat-reverse").prop("disabled", false);

                addMessage({
                    sender: "server",
                    message: "Disconnected"
                });
            };
            socket.onopen = function (e) {
                $( ".on-chat").prop("disabled", false);
                $( ".on-chat-reverse").prop("disabled", true);
                $( "input[name=message]", $(form) ).focus();
            };
            socket.onmessage = function( e ) {
                var event = JSON.parse( e.data );
                addMessage( event );
            };
        }
    }

    function adjustScroll() {
        var $chatroom = $("#chatroom");
        $chatroom.scrollTop( $chatroom.prop('scrollHeight') );
    }

    function addMessage( event ) {
        console.log(event);
        var $p = $( "<p/>" ).text( getNamecard( event.sender ) + event.message );
        $(".chatroom").append($p);
        adjustScroll();
    }

    function getNamecard( name ) {
        return "[" + name + "] ";
    }

    function disconnect() {
        socket.close();
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
        send: function() {
            send( $("input[name=message]", $(form) ).val() );
            $( "input[name=message]", $(form) ).val("");
        }
    }
});


var ws = sock( "#form-chat" );
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

});
