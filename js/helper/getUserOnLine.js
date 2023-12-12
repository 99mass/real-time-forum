import { linkApi } from "../helper/api_link.js";
// import { displayUsrOnLine  } from "../layout/corps.js";
let isThrottled = false;

// const userOnline = async () => {

//         const response = await fetch(`${linkApi}connectedUsers`, {
//             method: 'GET',
//             headers: {
//                 'Content-Type': 'application/json',
//             },
//         });

//         if (response.status === 200 && response.headers.get('content-type').includes('application/json')) {                       
//             // console.log( await response.json());
//             return await response.json();
            
//         }

// }

// function displayUsrOnLine(users) {

//     var users="";
// console.log(users);
//     for (let i = 0; i < users["Users"].length; i++) {
        
//         console.log(users[i]);
//         const user = users[i];    
//         if (user=="there's no user online") {
//         users=`<div class="user">  
//                 <div class="user-infos">                               
//                     <img src="assets/user-profile-svgrepo-com.svg" alt="">
//                     <div>
//                         <p><span>${user}</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
//                     </div>
//                 </div> 
//                 <div class="chat-text">
//                     <span>chat</span>
//                     <img src="assets/chat-dots-svgrepo-com.svg" alt="">
//                 </div>
//             </div>
//         `;
//     };
        
//     // }
//     return users;
    
// }

// const userOnline = () => {
//     const containUsers=document.querySelector('.bloc-users-on-line');
//     fetch(`${linkApi}connectedUsers`, {
//         method: 'GET',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//     })
//     .then(response => {
//         if (response.status === 200 && response.headers.get('content-type').includes('application/json')) {                       
//             return response.json();
//         }
//     })
//     .then(data => {
       
//         var users="";
//         if (data["Users"]) {                
//             for (let i = 0; i < data["Users"].length; i++) {
            
//                 console.log(data["Users"][i]);
//                 const user = data["Users"][i];   
//                 users+=`<div class="user">  
//                         <div class="user-infos">                               
//                             <img src="assets/user-profile-svgrepo-com.svg" alt="">
//                             <div>
//                                 <p><span>${user}</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
//                             </div>
//                         </div> 
//                         <div class="chat-text">
//                             <span>chat</span>
//                             <img src="assets/chat-dots-svgrepo-com.svg" alt="">
//                         </div>
//                     </div>
//                 `;                
//             }
//         }else{
//             users=`<div class="no-user">  ${data["message"]}</div>`;
//         }
//         containUsers.innerHTML=users;
//     })
//     .catch(error => {
//         console.error('Error:', error);
//     });
// }

// export { userOnline }