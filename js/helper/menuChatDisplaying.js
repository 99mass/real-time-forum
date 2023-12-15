
const chatContainerDisplaying = (chatText, userNameOnline, menuDots, chatContainer, UsernameinputChat) => {
    const chatBody = document.querySelectorAll('.chat-body');
    for (let i = 0; i < chatText.length; i++) {
        const btnChat = chatText[i];

        btnChat.addEventListener('click', () => {
            for (let j = 0; j < chatBody.length; j++) {
                chatBody[j].style.display = "none";
            }

            chatBody[i].style.display = "block";
            chatContainer.style.display = "block";
            UsernameinputChat.value = userNameOnline[i].textContent.trim();
        });
    }

    menuDots.addEventListener('click', () => {
        chatContainer.style.display = "none";
        for (let i = 0; i < chatText.length; i++) {
            chatBody[i].style.display = "none";
        }
    });
}

const alertMessage=()=>{
    return `
            <div class="alert">
            <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span> 
            <strong>Info!</strong> you have a new message.
        </div>
    `;
}

export{chatContainerDisplaying}