
var datas = [];
var columns = [];

var searchParams = new URLSearchParams(window.location.search)
//searchParams.has('action') // true
var date_r = searchParams.get('action')


function enter(){
   

     
    $.ajax({
     type: "GET",
     crossDomain: true,
     cache: false,
     data:{
         action:date_r,
     },
     url: '/attendence/gethistory',
     dataType: "json",
     async: false,
     success: function(result){
             datas=result.datas;
             columns=result.column;
             console.log(datas);
             console.log(columns);

             //alert("success.");

         
     },
     error: function(jqXHR, textStatus, errorThrown) {
         console.log(jqXHR);
         console.log(textStatus);
         console.log(errorThrown);
     }
 });
     

 

}


$(document).ready(function() {  
    enter();
    var hidden = $.fn.dataTable.absoluteOrder( [
        { value: 'XXX', position: 'top' }
      ] );
  
      
    $('#att_table').empty();
    console.log(datas);
    console.log(columns);
    $('#att_table').DataTable({
        // "serverSide":true,
        // "JQueryUI": true,
       "processing": true,
        "columns":columns,  //第三欄使用data中的startdate,
        "data":datas,
        "iDisplayLength": 10,// 每頁顯示筆數 
        "fixedColumns":   {
            leftColumns: 5,
        },
        "columnDefs": [{
          //targets: -1,
          //visible: false,
          type: hidden
        }],
        "scrollX": "200px",
       // "scrollY": "600px",

    });  
   
}

);  
