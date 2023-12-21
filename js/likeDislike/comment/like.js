import { linkApi } from "../../helper/api_link.js";
import { isGoodNumber } from "../../helper/utils.js";




const liskeComment=(likeComment,likeCommentId,likeCommentScore,dislikeCommentScore,dislikeComment)=>{
   
    for (let i = 0; i < likeComment.length; i++) {
        const btnlikeComment = likeComment[i];

        const handleClick = async function() {
            let commentID=likeCommentId[i].textContent.trim();
            try {
                const response = await fetch(`${linkApi}likecomment`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({CommentID: commentID}),

                });

                if(response.status === 200){

                   let countLike=isGoodNumber(likeCommentScore[i].textContent.trim());
                   let countDisLike=isGoodNumber(dislikeCommentScore[i].textContent.trim());
                    if (countDisLike>0) countDisLike-=1;
                   if (countLike===null) return ;
                   
                    if (!btnlikeComment.className.includes('like-post-color')) {                    
                        likeCommentScore[i].textContent=countLike+1;
                        btnlikeComment.style.color="darkgoldenrod";
                        likeCommentScore[i].style.color="darkgoldenrod";

                        dislikeCommentScore[i].textContent=countDisLike;
                        dislikeComment[i].style.color="#ffffff";
                        dislikeCommentScore[i].style.color="#ffffff";

                        btnlikeComment.classList.add('like-post-color');
                        dislikeComment[i].classList.remove('dislike-post-color');
                    }

                }else{

                    const er = await response.json();
                    console.error('Error : ' + er.message);
                    return;
                }
                
            } catch (error) {
            console.error("error : "+error);
            }
        };

        btnlikeComment.addEventListener('click', handleClick);
        
    }
}


export{
    liskeComment
}