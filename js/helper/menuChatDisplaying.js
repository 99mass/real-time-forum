
const chatContainerDisplaying=(chatText,userNameOnline,menuDots,chatContainer,UsernameinputChat)=>{
     for (let i = 0; i < chatText.length; i++) {
        const btnChat = chatText[i];
        btnChat.addEventListener('click',()=>{         
            chatContainer.style.display="block";
            UsernameinputChat.value=userNameOnline[i].textContent.trim();
            console.log(userNameOnline[i].textContent.trim());
            return;        
        });        
     }

     menuDots.addEventListener('click',()=>{
        chatContainer.style.display="none";
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