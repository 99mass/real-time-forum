
const userOnline = (data) => {

            const containUsers=document.querySelector('.bloc-users-on-line');
            var users="";
            
            if (data["message"]) {    
                users=`<div class="no-user">  ${data["message"]}</div>`;                    
            }else{
                for (let i = 0; i < data.length; i++) {
                    const user = data[i];   
                    users+=`<div class="user">  
                            <div class="user-infos">                               
                                <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                <div>
                                    <p><span>${user}</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                </div>
                            </div> 
                            <div class="chat-text">
                                <span>chat</span>
                                <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                            </div>
                        </div>
                    `;                
                }
            }
            containUsers.innerHTML=users;

}

export { userOnline }