import { linkApi } from "../helper/api_link.js";

// const sendMessage=()=>{
//     var socketMessage = new WebSocket("ws://localhost:8080/message");

//     // send user connected
//     socketMessage.onopen = () => {
//         socketMessage.send(JSON.stringify({ Username: data["User"]["Username"]}));
//         console.log("main");
//     }
//     const formChat=document.querySelector('.form-chat');
//     formChat.addEventListener('submit', async function (event) {
//         event.preventDefault();


//         let _Sender = document.querySelector('input[name="Sender"]').value.trim();
//         let _Recipient = document.querySelector('input[name="Recipient"]').value.trim();
//         let _Message = document.querySelector('textarea[name="Message"]').value.trim();

//         let data = {
//             Sender: _Sender,
//             Recipient: _Recipient,
//             Message: _Message
//         };  
//         // console.log(data); 
//         // Créez la connexion WebSocket 
//         var socket = new WebSocket("ws://localhost:8080/message");
//         socket.onopen = () => {
//             socket.send(JSON.stringify({ Username: _Sender}));
//             socket.send(JSON.stringify(data));
//             console.log("ok");
//         }


//     });

//     socketMessage.onmessage = (message) => {
//         var _data = JSON.parse(message.data);
//         console.log(_data);
//     }
// }


const sendMessage = (userName_) => {
    // Create WebSocket connection
    const socket = new WebSocket("ws://localhost:8080/message");

    // Define data
    // const data = {
    //     User: {
    //         Username: userName_ // Replace with actual username
    //     }
    // };

    // Send user connected message when connection is open
    socket.onopen = () => {
        socket.send(JSON.stringify({ Username: userName_}));
        console.log("main");
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
            socket.send(JSON.stringify({ Username: _Sender}));
            socket.send(JSON.stringify(messageData));
            console.log("ok");
        }
    });

    // Handle incoming messages
    socket.onmessage = (message) => {
        var _data = JSON.parse(message.data);
        console.log(_data);
    };

    // Handle errors
    socket.onerror = (error) => {
        console.error(`WebSocket error: ${error}`);
    };
}

export{sendMessage}