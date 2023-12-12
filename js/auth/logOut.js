import { linkApi } from "../helper/api_link.js";

const logOut=(_btn)=>{

    let _sessionId = getCookie('sessionID');
    _btn.addEventListener('click', async function(event) {
        try {
            const response = await fetch(`${linkApi}signout`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({session: _sessionId}),

            });

            if(response.status === 200){
            
                window.location.reload();
            }else{
                alert("a mistake is trying again");
            }
            
        } catch (error) {
           console.log("error : "+error);
        }
     })
}
function getCookie(name) {
    let cookieArr = document.cookie.split(";");
    for(let i = 0; i < cookieArr.length; i++) {
        let cookiePair = cookieArr[i].split("=");
        if(name == cookiePair[0].trim()) {
            return decodeURIComponent(cookiePair[1]);
        }
    }
    return null;
}

export {logOut}