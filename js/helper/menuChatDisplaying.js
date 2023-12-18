import { OldChatMessage } from "../chat/chatOldMessage.js";
import { sendMessages, recipientMessages } from "../layout/corps.js";
import { chatDateFormatter,timeAgo } from "../helper/utils.js";
import { throttle } from "../helper/utils.js";


const chatContainerDisplaying = (chatText, userNameOnline, menuDots, chatContainer, UsernameinputChat) => {
    const chatBody = document.querySelectorAll('.chat-body');
    const autherUser=document.querySelector(".auther-user");
    let messageQueue = [];

    for (let i = 0; i < chatText.length; i++) {
        const btnChat = chatText[i];

        // Create a new WebSocket connection for each user to get messsage
        const socket = new WebSocket("ws://localhost:8080/communication");

        btnChat.addEventListener('click', () => {
          
            // Send user connected message when connection is open
            socket.onopen = () => {
                console.log("WebSocket communication on.");
            }
            for (let j = 0; j < chatBody.length; j++) {
                chatBody[j].style.display = "none";
            }
            if (chatBody[i]) {
                 chatBody[i].style.display = "block";                
                 setTimeout(() => {
                     chatBody[i].scrollTo(0,chatBody[i].scrollHeight);
                }, 10);
            }
            chatContainer.style.display = "block";
            UsernameinputChat.value = userNameOnline[i].textContent.trim();
            autherUser.innerHTML=userNameOnline[i].textContent.trim();
            let _User1 = document.querySelector('.user').textContent.trim();
            let _User2 = userNameOnline[i].textContent.trim();

            let Messdata = {
                User1: _User1,
                User2: _User2,
            }
            console.log(Messdata);
            if (socket.readyState === WebSocket.OPEN) {
                socket.send(JSON.stringify(Messdata));
                console.log("message send by communication.");
            }

            socket.onmessage = (message) => {
                var _datas = JSON.parse(message.data);
                console.log(_datas);
            
                if (_datas) {
                   
                    messageQueue.push(..._datas);
                }
                if (messageQueue.length > 0) {
                    // Display the last 10 messages when the chat is opened
                    for (let j = 0; j < 10 && messageQueue.length > 0; j++) {
                        // Remove the last message from the queue
                        let _data = messageQueue.shift();
            
                        let formattedDate = timeAgo(_data["Created"]);
                        console.log(formattedDate);
            
                        let newMessage;
                        if (_data["Sender"] == _User1) {
                            newMessage = sendMessages(_data["Sender"], _data["Message"], formattedDate);
                        }
            
                        if (_data["Recipient"] == _User1) {
                            newMessage = recipientMessages(_data["Sender"], _data["Message"], formattedDate);
                        }
            
                        // Append new message to chatBody[i]
                        chatBody[i].appendChild(newMessage);
                    }
            
                    // Scroll to the bottom of chatBody[i]
                    chatBody[i].scrollTop = chatBody[i].scrollHeight;
                }
            };
            


            chatBody[i].addEventListener('scroll', throttle((event) => {
              
                // Check if user has scrolled to the top
                if (event.target.scrollTop === 0) {
                    // Check if there are any messages in the queue
                    if (messageQueue.length > 0) {
                        // Store current scroll height
                        let scrollHeight = chatBody[i].scrollHeight;
            
                        // Remove the last 3 messages from the queue and add them to chatBody[i]
                        for (let j = 0; j < 10; j++) {
                            if (messageQueue.length > 0) {
                                alert('ok')
                                let _data = messageQueue.shift();
            
                                let formattedDate = timeAgo(_data["Created"]);
            
                                let newMessage;
                                if (_data["Sender"] == _User1) {
                                    newMessage = sendMessages(_data["Sender"], _data["Message"], formattedDate);
                                }
            
                                if (_data["Recipient"] == _User1) {
                                    newMessage = recipientMessages(_data["Sender"], _data["Message"], formattedDate);
                                }
            
                                // Insert new message at the top of chatBody[i]
                                chatBody[i].insertBefore(newMessage, chatBody[i].firstChild);
                            }
                        }
            
                        // Adjust scroll position to prevent jumping
                        chatBody[i].scrollTop = chatBody[i].scrollHeight - scrollHeight;
                    }
                }
            }, 1000));

            // return;
        });
 
    }


    menuDots.addEventListener('click', () => {
        chatContainer.style.display = "none";
        for (let i = 0; i < chatText.length; i++) {
            if (chatBody[i])
                chatBody[i].style.display = "none";
        }
    });
}



const alertMessage = () => {
    return `
            <div class="alert">
            <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span> 
            <strong>Info!</strong> you have a new message.
        </div>
    `;
}

export { chatContainerDisplaying }