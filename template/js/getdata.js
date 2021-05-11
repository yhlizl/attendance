



function enter(){
   
     
         $.ajax({
          type: "POST",
          crossDomain: true,
          cache: false,
          data:{
          
          },
          url: '/attendence/gettoday',
          dataType: "json",
          success: function(result){
              console.log(result['NAME']);
             
          
                  console.log("success");
                  //alert("success.");
         
              
          },
          error: function(jqXHR, textStatus, errorThrown) {
              console.log(jqXHR);
              console.log(textStatus);
              console.log(errorThrown);
          }
      });
          
      
     
}

enter();