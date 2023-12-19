
const userOnline = (data,tab) => {
    console.log(tab);
    const containUsers = document.querySelector('.bloc-users-on-line');
    var users = "";

    if (data["message"]) {
        users = `<div class="no-user">  ${data["message"]}</div>`;
    } else {
        for (let i = 0; i < data.length; i++) {
            const user = data[i];
            let status = user['Status'] == "online" ? 'status-active-svgrepo-com.svg' : 'status-no-active-svgrepo-com.svg';
            if (user['Username']==tab[i].Sender) {
                let t=`<span>${tab[i].NumberMessage}</span>`
                console.log(t);
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
                            <div class="notification">
                                <img src="assets/notification-bell-svgrepo-com.svg" alt="">
                              
                            </div>
                            <div class="chat-text btn-chat ">
                                <span>chat</span>
                                <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                            </div>
                        </div>
                    `;
        }
    }
    containUsers.innerHTML = users;

}

export { userOnline }