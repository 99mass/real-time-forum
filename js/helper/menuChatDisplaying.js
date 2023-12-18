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
                console.log(Messdata);
                console.log("message send via communication.");
            }

            socket.onmessage = (message) => {
                var _datas = JSON.parse(message.data);
                console.log(_datas);
            
                if (_datas) {
                    messageQueue.push(..._datas);
                }
                if (messageQueue.length < 10) {
                    for (let k = 0 ; k <  messageQueue.length; k++) {
                        if (messageQueue[k]) {                                                   
                            let _data = messageQueue[k];            
                            let formattedDate = timeAgo(_data["Created"]);                                   
                            let newMessage;

                            if (_data["Sender"] == _User1) {
                                newMessage = sendMessages(_data["Sender"], _data["Message"], formattedDate);
                            }
                
                            if (_data["Recipient"] == _User1) {
                                newMessage = recipientMessages(_data["Sender"], _data["Message"], formattedDate);
                            }
                                    
                            chatBody[i].appendChild(newMessage);
                        }
                            
                        chatBody[i].scrollTop = chatBody[i].scrollHeight;
                     }
                }else {
                    var tempMessageQueue = [];
                    // Get the last 10 messages
                    for (let t = 0; t < Math.max(10, messageQueue.length); t++) {
                        if (messageQueue[t]) {                            
                            let tTemp = messageQueue.pop();
                            tempMessageQueue.push(tTemp);
                        }
                    }
                    if (messageQueue[0]) {                        
                        tempMessageQueue.push(messageQueue.pop());
                    }
                    // Display the messages in reverse order
                    for (let j = tempMessageQueue.length-1 ; j >= 0; j--) {
                        if (tempMessageQueue[j]) {                                                   
                            let _data = tempMessageQueue[j];            
                            let formattedDate = timeAgo(_data["Created"]);                                   
                            let newMessage;

                            if (_data["Sender"] == _User1) {
                                newMessage = sendMessages(_data["Sender"], _data["Message"], formattedDate);
                            }
                
                            if (_data["Recipient"] == _User1) {
                                newMessage = recipientMessages(_data["Sender"], _data["Message"], formattedDate);
                            }
                                    
                            chatBody[i].appendChild(newMessage);
                        }
                            
                        chatBody[i].scrollTop = chatBody[i].scrollHeight;
                     }
                }
            };

            chatBody[i].addEventListener('scroll', throttle((event) => {
                // Check if user has scrolled to the top
                if (event.target.scrollTop === 0) {
                    let scrollHeightBefore = chatBody[i].scrollHeight;
            
                    for (let t = 0; t < Math.max(10, messageQueue.length-1); t++) {
                        let _data = messageQueue.pop();
                        if (_data) {                          
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
                    let scrollHeightAfter = chatBody[i].scrollHeight;
                    chatBody[i].scrollTop = chatBody[i].scrollTop + (scrollHeightAfter - scrollHeightBefore);
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