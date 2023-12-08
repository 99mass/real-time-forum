
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
const rigthtBloc=(_rigthtBloc)=>{
    _rigthtBloc.innerHTML=`<h2> <img src="assets/right-arrow-svgrepo-com.svg" alt=""><span>Users On line</span> </h2>
                                <div class="bloc">
                                <div class="user">  
                                    <div class="user-infos">                               
                                        <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                        <div>
                                            <p><span>mass</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                        </div>
                                    </div> 
                                    <div class="chat-text">
                                        <span>chat</span>
                                        <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                                    </div>
                                </div>
                                <div class="user">  
                                    <div class="user-infos">                               
                                        <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                        <div>
                                            <p><span>abdou</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                        </div>
                                    </div> 
                                    <div class="chat-text">
                                        <span>chat</span>
                                        <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                                    </div>
                                </div>
                                <div class="user">  
                                    <div class="user-infos">                               
                                        <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                        <div>
                                            <p><span>adama</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                        </div>
                                    </div> 
                                    <div class="chat-text">
                                        <span>chat</span>
                                        <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                                    </div>
                                </div>
                                <div class="user">  
                                    <div class="user-infos">                               
                                        <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                        <div>
                                            <p><span>yassine</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>                            
                                        </div>
                                    </div> 
                                    <div class="chat-text">
                                        <span>chat</span>
                                        <img src="assets/chat-dots-svgrepo-com.svg" alt="">
                                    </div>
                                </div>
                                
                            </div>    
    `;
}

function createPostMenue() {

    let _div=  document.createElement('div');
    _div.className="create-post-menue";
    _div.innerHTML=`<span>
                        <img src="assets/postcard-svgrepo-com.svg" alt="">
                        <span>My Posts</span>
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

function posts(_posts) {
    var posts="",lastPost="",lastFormComment="",lastBlocComment="";
    for (let i = 0; i < _posts.length; i++) {
        const post = _posts[i];
        const title=post["Posts"]["Title"];
        const content=post["Posts"]["Content"];
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
                                    <p><span>${Username}</span><img  src="assets/status-active-svgrepo-com.svg" alt=""> </p>
                                    <p>3 weeks ago</p>
                                </div>
                            </div>
                            <div class="like-comment-block">
                                <div class=" ${classColorDisliked} dislike-post"> <span class="scoreDisLike">${postDislikCount}</span> dislikes <span class="id-post-dislike" style="display: none;"> ${postId}</span></div>
                                <div class=" ${classColorLiked } like-post"><span class="scoreLike">${postLikeCount}</span> likes<span class="id-post-like"  style="display: none;"> ${postId}</span></div>
                                <div class="comment">${commentCount} comments <span class="Id-post" style="display:none">${postId}</span> </div>
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
                const ComCreatedAt= comment["Comment"]["CreatedAt"];
                const ComPostId= comment["Comment"]["PostID"];
                idPost=ComPostId;
                const ComUserId= comment["Comment"]["UserID"];
                const ComLike=comment["CommentLike"];
                const ComDislike=comment["CommentDislike"];
                const ComUserName=comment["User"]["Username"];             


                _comments+=`<div >
                            <div class="comment-text "><pre class="card-description">${ComContent}</pre></div>
                            <div class="content-comment-like">
                            <div class="content-comment">
                                <div class="comment">                                  
                                    <img src="assets/user-profile-svgrepo-com.svg" alt="">
                                    <div>
                                        <p><span>${ComUserName}</span> </p>
                                        <p>2 weeks ago</p>
                                    </div>
                                </div>
                                <div class="like-comment-block">
                                    <div>${ComDislike} dislikes</div>
                                    <div>${ComLike} Likes</div>
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





export {leftBloc,middleBloc,rigthtBloc,displayComment}