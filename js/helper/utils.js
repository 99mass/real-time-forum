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
      return parseInt(secondsPast) + ' seconds ago';
  }

  if(secondsPast < 3600) {
      return parseInt(secondsPast/60) + ' minutes ago';
  }

  if(secondsPast <= 86400) {
      return parseInt(secondsPast/3600) + ' hours ago';
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

const commentTemporel=(ComContent,ComUserName)=>{
    let currentDate = new Date();
    let formattedDate = currentDate.getFullYear() + '-' + 
        String(currentDate.getMonth() + 1).padStart(2, '0') + '-' + 
        String(currentDate.getDate()).padStart(2, '0') + ' ' + 
        String(currentDate.getHours()).padStart(2, '0') + ':' + 
        String(currentDate.getMinutes()).padStart(2, '0') + ':' + 
        String(currentDate.getSeconds()).padStart(2, '0');
    let date=timeAgo(formattedDate);
    return `<div >
            <div class="comment-text "><pre class="card-description">${ComContent}</pre></div>
            <div class="content-comment-like">
            <div class="content-comment">
                <div class="comment">                                  
                    <img src="assets/user-profile-svgrepo-com.svg" alt="">
                    <div>
                        <p><span>${ComUserName}</span> </p>
                        <p>${date}</p>
                    </div>
                </div>
                <div class="like-comment-block">
                    <div class=" dislike-comment"> <span class="scorecommentDisLike">${0}</span> dislikes <span class="id-comment-dislike" style="display: none;"> ${0}</span></div>
                    <div class=" like-comment"><span class="scorecommentLike">${0}</span> likes<span class="id-comment-like"  style="display: none;"> ${0}</span></div>
                </div>
            </div>
        </div>
            </div>
        <hr>
    `;

};

const debounce = (callback, delay)=>{
    var timer;
    return function(){
        var args = arguments;
        var context = this;
        clearTimeout(timer);
        timer = setTimeout(function(){
            callback.apply(context, args);
        }, delay)
    }
}

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

export{isGoodNumber,timeAgo,commentTemporel,throttle}