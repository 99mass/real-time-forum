import { linkApi } from "../helper/api_link.js";

const sendMessage=()=>{
    const formChat=document.querySelector('.form-chat');
    formChat.addEventListener('submit', async function (event) {
        event.preventDefault();


        let _Sender = document.querySelector('input[name="Sender"]').value.trim();
        let _Recipient = document.querySelector('input[name="Recipient"]').value.trim();
        let _Message = document.querySelector('textarea[name="Message"]').value.trim();

        let data = {
            Sender: _Sender,
            Recipient: _Recipient,
            Message: _Message
        };  
        // console.log(data); 
        // Créez la connexion WebSocket 
        var socketMessage = new WebSocket("ws://localhost:8080/message");
        socketMessage.onopen = () => {
            socketMessage.send(JSON.stringify({ Username: _Sender}));
            console.log("_Sender: "+_Sender);
            socketMessage.send(JSON.stringify(data));
            console.log("okurrrrrrrrrrrrrrrrrrrrrrrrr");
        }
        // send user connected
            // socket.send(JSON.stringify(data));
            // console.log("okurrrrrrrrrrrrrrrrrrrrrrrrr");
     
        socketMessage.onmessage = (message) => {
            var _data = JSON.parse(message.data);
            console.log(_data);
        }

    });
}
export{sendMessage}