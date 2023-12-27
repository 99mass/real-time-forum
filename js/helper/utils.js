const isGoodNumber=(score)=>{
    let likecount = Number(score);
    if (isNaN(likecount)) {
        console.error("Error: The value must be a number");
        return null
      } else if (!Number.isInteger(likecount) || likecount<0) {
        console.error("Error: The value must be an integer");
        return null
    } 
     return likecount;

}


function timeAgo(dateString) {
  const date = new Date(dateString);
  const now = new Date();

  const secondsPast = (now.getTime() - date.getTime()) / 1000;

  if(secondsPast < 60) {
      return parseInt(secondsPast) + ' sec ago';
  }

  if(secondsPast < 3600) {
      return parseInt(secondsPast/60) + ' mn ago';
  }

  if(secondsPast <= 86400) {
      return parseInt(secondsPast/3600) + ' h ago';
  }

  if(secondsPast > 86400) {
      const daysPast = parseInt(secondsPast/86400);
      if (daysPast < 7) {
          return daysPast + ' days ago';
      } else if (daysPast < 30) {
          return parseInt(daysPast/7) + ' weeks ago';
      } else if (daysPast < 365) {
          return parseInt(daysPast/30) + ' months ago';
      } else {
          return parseInt(daysPast/365) + ' years ago';
      }
  }
}


const commentTemporel=(ComContent,ComUserName,scoreDisLike,scoreLike,comId ,dateCom)=>{

    let formattedDate=timeAgo(dateCom)
    return `<div >
            <div class="comment-text "><pre class="card-description">${ComContent}</pre></div>
            <div class="content-comment-like">
            <div class="content-comment">
                <div class="comment">                                  
                    <img src="assets/user-profile-svgrepo-com.svg" alt="">
                    <div>
                        <p><span>${ComUserName}</span> </p>
                        <p>${formattedDate}</p>
                    </div>
                </div>
                <div class="like-comment-block">
                    <div class=" dislike-comment"> <span class="scorecommentDisLike">${scoreDisLike}</span> dislikes <span class="id-comment-dislike" style="display: none;">${comId}</span></div>
                    <div class=" like-comment"><span class="scorecommentLike">${scoreLike}</span> likes<span class="id-comment-like"  style="display: none;">${comId}</span></div>
                </div>
            </div>
        </div>
            </div>
        <hr>
    `;

};

const throttle = (callback, delay)=> {
    var last;
    var timer;
    return function () {
        var context = this;
        var now = +new Date();
        var args = arguments;
        if (last && now < last + delay) {
            // le délai n'est pas écoulé on reset le timer
            clearTimeout(timer);
            timer = setTimeout(function () {
                last = now;
                callback.apply(context, args);
            }, delay);
        } else {
            last = now;
            callback.apply(context, args);
        }
    };
}
const statusPostUser=(dataSorted)=>{
    
    const postUsername=document.querySelectorAll('.post-username');
    const postUsernameImg=document.querySelectorAll('.post-username-img');
    for (let i = 0; i < postUsername.length; i++) {
            for (let j = 0; j < dataSorted.length; j++) {
                let name=dataSorted[j]["Username"];
                let status=dataSorted[j]["Status"];
                if (name===postUsername[i].textContent.trim() && status=='offline') {
                    postUsernameImg[i].src="assets/status-no-active-svgrepo-com.svg";
                }
                if (name===postUsername[i].textContent.trim() && status=='online') {
                    postUsernameImg[i].src="assets/status-active-svgrepo-com.svg";
                }
            }
    }

}
const statusPostFilteredUser=()=>{
    const userNameOnline=document.querySelectorAll('.user-infos div .user-name-online'); 
    const _userNameOnlineImg=document.querySelectorAll('.user-infos div .user-name-online-img'); 
    const postUsername=document.querySelectorAll('.content-poster-like .content-poster .poster div p span');
    const postUsernameImg=document.querySelectorAll('.content-poster-like .content-poster .poster div p img');

    for (let i = 0; i < postUsername.length; i++) {
            for (let j = 0; j < userNameOnline.length; j++) {
                let name=userNameOnline[j].textContent.trim();
                let srcValue=_userNameOnlineImg[j].src;
                if (name==postUsername[i].textContent.trim() && srcValue.includes("no-active") ) {
                    postUsernameImg[i].src="assets/status-no-active-svgrepo-com.svg";
                }
                if (name==postUsername[i].textContent.trim() && !srcValue.includes("no-active") ) {
                    postUsernameImg[i].src="assets/status-active-svgrepo-com.svg";
                }
            }
    }
}

const alertMessage = (sender, Recipient) => {
    const div = document.createElement('div');
    div.className = 'notif';

    const span = document.createElement('span');
    span.className = 'welcome_text';

    span.innerHTML = `Hey, <span class="notif-user"> ${Recipient}</span>  you have a new message from <span class="notif-user"> ${sender}</span>`;

    div.appendChild(span);

    return div;
}

const loader=()=>{
    return `<div class="loader ">
                <div class="bar1"></div>
                <div class="bar2"></div>
                <div class="bar3"></div>
                <div class="bar4"></div>
                <div class="bar5"></div>
                <div class="bar6"></div>
                <div class="bar7"></div>
                <div class="bar8"></div>
                <div class="bar9"></div>
                <div class="bar10"></div>
                <div class="bar11"></div>
                <div class="bar12"></div>
        </div>
    `;
}
const loaderMessage=()=>{
    return `<div class="center">
                <div class="wave"></div>
                <div class="wave"></div>
                <div class="wave"></div>
                <div class="wave"></div>
                <div class="wave"></div>
                <div class="wave"></div>
                <div class="wave"></div>
                <div class="wave"></div>
                <div class="wave"></div>
                <div class="wave"></div>
            </div>
    `;
}

export{isGoodNumber,timeAgo,commentTemporel,statusPostUser,statusPostFilteredUser,throttle,alertMessage,loader,loaderMessage}