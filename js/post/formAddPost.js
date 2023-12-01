

const formAddPost=(formCreatPost)=>{
    const defaultCategorie=["Education", "Sport", "Art", "Culture", "Religion"];
    formCreatPost.addEventListener('submit', function(e) {
        e.preventDefault();
        // const  errorPost=document.querySelector('.error-post');
        let selectedCategories = [];
        document.querySelectorAll('input[name="categories"]:checked').forEach((checkbox) => {
            if (!defaultCategorie.includes(checkbox.value)) {
                alert('Categorie(s) invalid');
                // errorPost.style.diplay="block";
                // errorPost.innerHTML='Categorie(s) invalid';
                return;
            }
            selectedCategories.push(checkbox.value);
        });

        if (selectedCategories.length==0) {
            alert('Categorie(s) are required');
            // errorPost.style.diplay="block";
            // errorPost.innerHTML='Categorie(s) are required';
                return;
        }

        let title = document.querySelector('input[name="Title"]').value.trim();
        let content = document.querySelector('textarea[name="Content"]').value.trim();
        let file = document.querySelector('input[type="file"]').files[0];
        
        if (title==="") {
            alert('Title is required');
            // errorPost.style.diplay="block";
            // errorPost.innerHTML='Title is required';
            return 
        }
        if (content==="") {
            alert('Content is required');
            // errorPost.style.diplay="block";
            // errorPost.innerHTML='Content is required';
            return 
        }
        if (title.length>100) {
            alert('Title length is over than 100 characters');
            // errorPost.style.diplay="block";
            // errorPost.innerHTML='Title length is over than 100 characters';
            return 
        }

        if (file) {
            let fileType = file.type;
            let fileSize = file.size / (1024 * 1024); // size in MB

            if (!(['image/jpeg', 'image/svg+xml', 'image/png', 'image/gif'].includes(fileType) && fileSize <= 20)) {
             alert('File is not valid')
                // errorPost.style.diplay="block";
                // errorPost.innerHTML='File is not valid';
                return;
            }
        }


        console.log('Selected Categories:', selectedCategories);
        console.log('Title:', title);
        console.log('Content:', content);
        console.log('File:', file);
    });


}

export{formAddPost}