import { timeAgo } from "../helper/utils.js";


const leftBloc=(_leftBloc,CatgoryArray)=>{
       _leftBloc.innerHTML=`<div> <img src="assets/category-list-solid-svgrepo-com.svg" alt="">Categories</div>
       `;
       if (CatgoryArray) {
          CatgoryArray.forEach(category => {
            let _div=  document.createElement('div');
            _div.className="contenCatId";
            let _span=  document.createElement('span');
            _span.className="categoryId";
            _span.textContent=category.ID;
            _span.style.display="none";

            _div.textContent=category.NameCategory;
            _div.appendChild(_span)
            
            _leftBloc.appendChild(_div);
          });
       }
}

const middleBloc=(_middleBloc,_posts)=>{
    _middleBloc.appendChild(createPostMenue());

    _middleBloc.appendChild(contentPostBloc(_posts));


}
const rigthtBloc=(_rigthtBloc,userName)=>{

    _rigthtBloc.innerHTML=`<h2> <img src="assets/right-arrow-svgrepo-com.svg" alt=""><span>Users On line</span> </h2>
                            <div class="bloc bloc-users-on-line"> </div>  
                            <div class="chat-container">
                            <div class="chat-header">
                                <div></div>
                                <div class="midle-header">
                                    <div></div>
                                    <div>${userName}</div>
                                </div>
                                <div class="right-header"><img class="menu-dots" src="../assets/menu-dots-svgrepo-com.svg"></div>
                            </div>

                            <div class="chat-body">
                                <div class="sender-bloc">
                                <div class="user-sender-bloc">
                                     <img src="../assets/user-profile-svgrepo-com.svg"/>
                                     <p>breukh</p>
                                </div>
                                    <div class="message-date">
                                        <p>Nice too meet you</p>
                                        <p>12:15</p>
                                    </div>
                                </div>
                                <div class="receiver-bloc">
                                <div class="message-date-2">
                                    <p>Nice too meet you too</p>
                                    <p>16:30</p>
                                </div>
                                <div class="user-receiver-bloc">
                                    <img src="../assets/user-profile-svgrepo-com-2.svg"/>
                                    <p>samba</p>
                                </div>
                                </div>
                            </div>

                            <div class="btn-group-chat">
                                <input type="hidden" name="Username" class="Username-input-chat" />
                                <textarea name="message" class="chat-text" placeholder="chat here..." ></textarea>
                                <button type="button" class="btn-chat"> </button>
                            </div>

                            </div>  
    `;
}



function createPostMenue() {

    let _div=  document.createElement('div');
    _div.className="create-post-menue";
    _div.innerHTML=`<span>
                        <img src="assets/postcard-svgrepo-com.svg" alt="">
                        <span></span>
                    </span>
                    <button class="create-post-btn">Create Post</button>
    `;
    return _div;
}



function contentPostBloc(_posts) {
  
  let _div=  document.createElement('div');
    _div.className="content-post-block";
    _div.innerHTML=posts(_posts);
    return _div;
}

// function posts(_posts, start, end) {
//     var posts="",lastPost="",lastFormComment="",lastBlocComment="";
//     for (let i = start; i < end; i++) {
//         const post = _posts[i];

//         const title=post["Posts"]["Title"];
//         const content=post["Posts"]["Content"];
//         const CreatedAtPost=timeAgo(post["Posts"]["CreatedAt"]);
//         const Categories=post["Posts"]["Categories"];
//         const Username=post["User"]["Username"];
//         const Image=post["Posts"]["Image"];
//         const commentCount=post["CommentCount"];
//         const postLikeCount=post["PostLike"];
//         const postDislikCount=post["PostDislike"];
//         const userId=post["Posts"]["UserID"];
//         const postId=post["Posts"]["ID"];
//         const stateLiked=post["Liked"];
//         const stateDisLiked=post["Disliked"];
//         let classColorLiked=stateLiked ? "like-post-color" : "";
//         let classColorDisliked=stateDisLiked  ? "dislike-post-color" : "";



//         if (i===_posts.length-1 ){
//             lastPost="last-post";
//             lastFormComment="last-form-comment";
//             lastBlocComment="last-bloc-comment";
//         }

//         // recupere les categorie du post
//         let NameCategories="";
//         for (let j = 0; j < Categories.length; j++) {
//             const NameCategory = Categories[j]["NameCategory"];
//             NameCategories+=`<span>${NameCategory}</span>`;
//         }
        
//         let image=Image ? `<div class="image-post"><img src="/static/image/${Image}" alt=""></div>` : "";
        
//      posts += `<div class="one-post-block ${lastPost}">
//                 ${image}
//                 <div class="post-content">
//                     <h2>${title}</h2>
//                     <div class="post-text "><pre class="card-description">${content}</pre></div>
//                    <button class="myBtn">Read more</button>
//                     <div class="categorie-post">${NameCategories}</div>
//                     <div class="content-poster-like">
//                         <div class="content-poster">
//                             <div class="poster">                                  
//                                 <img src="assets/user-profile-svgrepo-com.svg" alt="">
//                                 <div>
//                                     <p><span>${Username}</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>
//                                     <p>${CreatedAtPost}</p>
//                                 </div>
//                             </div>
//                             <div class="like-comment-block">
//                                 <div class=" ${classColorDisliked} dislike-post"> <span class="scoreDisLike">${postDislikCount}</span> dislikes <span class="id-post-dislike" style="display: none;"> ${postId}</span></div>
//                                 <div class=" ${classColorLiked } like-post"><span class="scoreLike">${postLikeCount}</span> likes<span class="id-post-like"  style="display: none;"> ${postId}</span></div>
//                                 <div class="comment"><span class="commentCount">${commentCount} </span> comments <span class="Id-post" style="display:none">${postId}</span> </div>
//                             </div>
//                         </div>
//                     </div>
                
//                 </div>
//             </div> 
//          <div class="create-comment ${lastFormComment}"  style="display: none;">
//          <form class="form-comment" method="post">
//             <input type="hidden" name ="userId" value="${userId}" />
//             <input type="hidden" name ="postId" value="${postId}" />
//             <textarea type="text" name="ContentComment" id="Content" placeholder="Comment here..."  ></textarea>  
//             <input type="hidden" name="userName" value="${Username}" />

//             <button type="submit"> <img src="assets/send-svgrepo-com.svg" alt=""> </button>
//          </form>
//          </div>
//          <div class="bloc-comment ${lastBlocComment}" style="display: none;" > ${postId}</div>
//     ` ;
//     NameCategories="";
//  }

//     return posts
// }


function posts(_posts) {
    var posts="",lastPost="",lastFormComment="",lastBlocComment="";
    for (let i = 0; i < _posts.length; i++) {
        const post = _posts[i];

        const title=post["Posts"]["Title"];
        const content=post["Posts"]["Content"];
        const CreatedAtPost=timeAgo(post["Posts"]["CreatedAt"]);
        const Categories=post["Posts"]["Categories"];
        const Username=post["User"]["Username"];
        const Image=post["Posts"]["Image"];
        const commentCount=post["CommentCount"];
        const postLikeCount=post["PostLike"];
        const postDislikCount=post["PostDislike"];
        const userId=post["Posts"]["UserID"];
        const postId=post["Posts"]["ID"];
        const stateLiked=post["Liked"];
        const stateDisLiked=post["Disliked"];
        let classColorLiked=stateLiked ? "like-post-color" : "";
        let classColorDisliked=stateDisLiked  ? "dislike-post-color" : "";



        if (i===_posts.length-1 ){
            lastPost="last-post";
            lastFormComment="last-form-comment";
            lastBlocComment="last-bloc-comment";
        }

        // recupere les categorie du post
        let NameCategories="";
        for (let j = 0; j < Categories.length; j++) {
            const NameCategory = Categories[j]["NameCategory"];
            NameCategories+=`<span>${NameCategory}</span>`;
        }
        
        let image=Image ? `<div class="image-post"><img src="/static/image/${Image}" alt=""></div>` : "";
        
     posts += `<div class="one-post-block ${lastPost}">
                ${image}
                <div class="post-content">
                    <h2>${title}</h2>
                    <div class="post-text "><pre class="card-description">${content}</pre></div>
                   <button class="myBtn">Read more</button>
                    <div class="categorie-post">${NameCategories}</div>
                    <div class="content-poster-like">
                        <div class="content-poster">
                            <div class="poster">                                  
                                <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                <div>
                                    <p><span class="post-username">${Username}</span><img class="post-username-img"  src="assets/status-active-svgrepo-com.svg" alt=""> </p>
                                    <p>${CreatedAtPost}</p>
                                </div>
                            </div>
                            <div class="like-comment-block">
                                <div class=" ${classColorDisliked} dislike-post"> <span class="scoreDisLike">${postDislikCount}</span> dislikes <span class="id-post-dislike" style="display: none;"> ${postId}</span></div>
                                <div class=" ${classColorLiked } like-post"><span class="scoreLike">${postLikeCount}</span> likes<span class="id-post-like"  style="display: none;"> ${postId}</span></div>
                                <div class="comment"><span class="commentCount">${commentCount} </span> comments <span class="Id-post" style="display:none">${postId}</span> </div>
                            </div>
                        </div>
                    </div>
                
                </div>
            </div> 
         <div class="create-comment ${lastFormComment}"  style="display: none;">
         <form class="form-comment" method="post">
            <input type="hidden" name ="userId" value="${userId}" />
            <input type="hidden" name ="postId" value="${postId}" />
            <textarea type="text" name="ContentComment" id="Content" placeholder="Comment here..."  ></textarea>  
            <input type="hidden" name="userName" value="${Username}" />

            <button type="submit"> <img src="assets/send-svgrepo-com.svg" alt=""> </button>
         </form>
         </div>
         <div class="bloc-comment ${lastBlocComment}" style="display: none;" > ${postId}</div>
    ` ;
    NameCategories="";
 }

    return posts
}

function displayComment(bloComment ,comments,createCommentForm) {
    if (comments.length>0) {
    for (let c = 0; c < bloComment.length; c++) {
        var _comments="";
            const bloc=bloComment[c];
            let idPost    
            var  _comments="";
           
                
            for (let i = 0; i < comments.length; i++) {
                const comment = comments[i];      
                const ComContent= comment["Comment"]["Content"];
                const ComCreatedAt=timeAgo(comment["Comment"]["CreatedAt"]);
                const ComPostId= comment["Comment"]["PostID"];
                const ComId= comment["Comment"]["ID"];

                idPost=ComPostId;
                const ComUserId= comment["Comment"]["UserID"];
                const ComLike=comment["CommentLike"];
                const ComDislike=comment["CommentDislike"];
                const ComUserName=comment["User"]["Username"];             
                const stateLiked=comment["Liked"];
                const stateDisLiked=comment["Disliked"];
                let classColorLiked=stateLiked ? "like-post-color" : "";
                let classColorDisliked=stateDisLiked  ? "dislike-post-color" : "";

                _comments+=`<div >
                            <div class="comment-text "><pre class="card-description">${ComContent}</pre></div>
                            <div class="content-comment-like">
                            <div class="content-comment">
                                <div class="comment">                                  
                                    <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                    <div>
                                        <p><span>${ComUserName}</span> </p>
                                        <p>${ComCreatedAt}</p>
                                    </div>
                                </div>
                                <div class="like-comment-block">
                                    <div class=" ${classColorDisliked} dislike-comment"> <span class="scorecommentDisLike">${ComDislike}</span> dislikes <span class="id-comment-dislike" style="display: none;"> ${ComId}</span></div>
                                     <div class=" ${classColorLiked } like-comment"><span class="scorecommentLike">${ComLike}</span> likes<span class="id-comment-like"  style="display: none;"> ${ComId}</span></div>
                                </div>
                            </div>
                        </div>
                            </div>
                        <hr>
                `;
            }   
            if (idPost==bloc.textContent.trim()) {
                if (bloc.style.display==="none") {
                    createCommentForm[c].style.display="block";
                    bloc.style.display="block";
                    bloc.innerHTML=  _comments ;
                }
            }else{
                if (bloc.style.display==="block") {                    
                    bloc.textContent=  idPost ;
                    createCommentForm[c].style.display="none";
                    bloc.style.display="none";
                }
            }
            
        }       
        
    }
    
}





export {leftBloc,middleBloc,rigthtBloc,displayComment,createPostMenue,posts}