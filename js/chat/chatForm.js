import { sendMessages, recipientMessages } from "../layout/corps.js";
import {  timeAgo,alertMessage ,songNotification} from "../helper/utils.js";


const sendMessage = (userName_) => {
    // Create WebSocket connection
    const socket = new WebSocket("ws://localhost:8080/message");

    // Send user connected message when connection is open
    socket.onopen = () => {
        socket.send(JSON.stringify({ Username: userName_ }));
    };

    // Handle form submission
    const formChat = document.querySelector('.form-chat');
    formChat.addEventListener('submit', async function (event) {
        event.preventDefault();

        let _Sender = document.querySelector('input[name="Sender"]').value.trim();
        let _Recipient = document.querySelector('input[name="Recipient"]').value.trim();
        let _Message = document.querySelector('textarea[name="Message"]').value.trim();

        let messageData = {
            Sender: _Sender,
            Recipient: _Recipient,
            Message: _Message
        };

        // Send message when connection is open
        if (socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify(messageData));
            document.querySelector('textarea[name="Message"]').value = "";
        }
    });

    // Handle incoming messages
    socket.onmessage = (message) => {
        var _data = JSON.parse(message.data);
        let formattedDate = timeAgo(_data["created"]);

        // Sender
        if (_data["sender"] == userName_) {          
          let  noChatTex=  document.querySelector('.no-chat');
          if (noChatTex) noChatTex.style.display='none';
            var _recipient = document.querySelector(`.chat-container .chat-body-container .${_data["recipient"]}`);
            let send = sendMessages(_data["sender"], _data["message"],formattedDate);
            _recipient.appendChild(send);
        }

        // recipient
        if (_data["recipient"] == userName_) {
            var _sender = document.querySelector(`.chat-container .chat-body-container .${_data["sender"]}`);
            let recip = recipientMessages(_data["sender"], _data["message"], formattedDate);
            _sender.appendChild(recip);

            //   paly song notification
            let audio = songNotification();
            audio.play();

            document.querySelector('body').appendChild(alertMessage(_data["sender"],_data["recipient"]))
            setTimeout(() => {            
                document.querySelector('.notif').remove();
            }, 7000);
        }
     
    };

    // Handle errors
    socket.onerror = (error) => {
        console.error(`WebSocket error: ${error}`);
    };
}

export { sendMessage }
