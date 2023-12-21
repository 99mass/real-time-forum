import { linkApi } from "../../helper/api_link.js";
import { isGoodNumber } from "../../helper/utils.js";


const DisLiskeComment=(dislikeComment,dislikeCommentId,likeCommentScore,dislikeCommentScore,likeComment)=>{
   
    for (let i = 0; i < dislikeComment.length; i++) {
        const btnDislikeComment = dislikeComment[i];

        const handleClick = async function() {

            let commentID=dislikeCommentId[i].textContent.trim();
            try {
                const response = await fetch(`${linkApi}dislikecomment`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({CommentID: commentID}),

                });

                if(response.status === 200){

                   let countDisLike =isGoodNumber( dislikeCommentScore[i].textContent.trim());
                   let countLike =isGoodNumber(likeCommentScore[i].textContent.trim());
                    if (countLike>0) countLike-=1;
                   if (countDisLike===null) return ;
                   
                    if (!btnDislikeComment.className.includes('dislike-post-color')) {                    
                        dislikeCommentScore[i].textContent=countDisLike+1;
                        btnDislikeComment.style.color="gray";
                        dislikeCommentScore[i].style.color="gray";

                        likeCommentScore[i].textContent=countLike;
                        likeCommentScore[i].style.color="#ffffff";
                        likeComment[i].style.color="#ffffff";
                        btnDislikeComment.classList.add('dislike-post-color');
                        likeComment[i].classList.remove('like-post-color');
                    }

                    return
                }else{

                    const er = await response.json();
                    console.error('Error : ' + er.message);
                    return;
                }
                
            } catch (error) {
            console.error("error : "+error);
            }
        };

        btnDislikeComment.addEventListener('click', handleClick);
        
    }
}

  export{DisLiskeComment
  }