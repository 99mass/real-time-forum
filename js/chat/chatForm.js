import { linkApi } from "../helper/api_link.js";

const sendMessage=(socketMessage)=>{
    const formChat=document.querySelector('.form-chat');
    formChat.addEventListener('submit', async function (event) {
        event.preventDefault();


        let _Sender = document.querySelector('input[name="Sender"]').value.trim();
        let _Recipient = document.querySelector('input[name="Recipient"]').value.trim();
        let _Message = document.querySelector('textarea[name="Message"]').value.trim();

        // if (_Message.lenght>2) {
        //     alert("message must not exceed 200 characters");
        //     return;
        // }
        // alert(_Message)

        let data = {
            Sender: _Sender,
            Recipient: _Recipient,
            Message: _Message
        };  
        // console.log(data); 
        // Créez la connexion WebSocket 
        var socketMessage = new WebSocket("ws://localhost:8080/message");

        // send user connected
        socketMessage.onopen = () => {
            socketMessage.send(JSON.stringify(data));
            console.log("ok");
        }
     
        socketMessage.onmessage = (message) => {
            var _data = JSON.parse(message.data);
            console.log(_data);
        }

    });
}
export{sendMessage}