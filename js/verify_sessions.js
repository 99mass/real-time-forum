const isSessionFound = () => {
    // let sessionId = sessionStorage.getItem('sessionID');
    let sessionId = getCookie('sessionID');
    if (sessionId) {
        console.log("Session found: " + sessionId);
        fetch('http://127.0.0.1:8080/verifySession', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({sessionId: sessionId}),
        })
        .then(response => response.json())
        .then(data => console.log(data))
        .catch((error) => {
          console.error('Error:', error);
        });
    } else {
        console.log("No session found");
    }
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

export {isSessionFound}