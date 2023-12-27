import {loader } from "../helper/utils.js";

const userOnline = (data,tab) => {

    const containUsers = document.querySelector('.bloc-users-on-line');
    containUsers.innerHTML=loader();
    containUsers.classList.add('temp-class');
    
    var users = "";

    if (data["message"]) {        
        users = `<div class="no-user">  ${data["message"]}</div>`;
    } else {
        for (let i = 0; i < data.length; i++) {
            let user = data[i];
            for (let j = 0; j < tab.length; j++) {
                let not = tab[j];     
                if (user['Username']===not.Sender ) {           
                    user.NumberMessage = not.NumberMessage;         
                }     
           
            }
        }

        for (let i = 0; i < data.length; i++) {
            let user = data[i];
            let status = user['Status'] == "online" ? 'status-active-svgrepo-com.svg' : 'status-no-active-svgrepo-com.svg';           

            let countNotif=""
          
            if ( user['NumberMessage']>0 ) {                    
                countNotif=` <img src="assets/notification-bell-svgrepo-com.svg" alt="">
                            <span class="notif-value-${user['Username']}" >${user['NumberMessage']}</span>                                              
                `;           
            }     
               
            users += `<div class="user user-on-line ">  
                            <div class="user-infos">                               
                                <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                
                                <div>                                        
                                    <p>
                                        <span class="user-name-online">${user['Username']}</span>                                            
                                        <img class="user-name-online-img"  src="assets/${status}" alt="">                                            
                                    </p>                            
                                </div>
                            </div> 
                            <div class="notification number-message-${user['Username']}" >
                                ${countNotif}
                             </div> 
                            <div  class="chat-text btn-chat ">
                                <span>chat</span>
                                <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                            </div>
                        </div>
                    `;
        }
        
    }
   
   setTimeout(() => {   
        let loaders= document.querySelector('.loader'); 
        containUsers.classList.remove('temp-class');
       if (loaders) loaders.remove();
       containUsers.innerHTML = users;
   }, 1000);
}

export { userOnline }