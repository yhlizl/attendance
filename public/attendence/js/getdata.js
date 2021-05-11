



function enter(){
   
     
         $.ajax({
          type: "GET",
          crossDomain: true,
          cache: false,
          data:{
            
          },
          url: '/attendence/gettoday',
          dataType: "json",
          success: function(result){
              
             
          
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




function edit(date_s,empid,remark){
   
    console.log(date_s,empid,remark);
    $.ajax({
     type: "POST",
     crossDomain: true,
     cache: false,
     data:{
       remark:remark,
       date_s:date_s,
       empid:empid,
     },
     url: '/attendence/edit',
     dataType: "json",
     success: function(result){
         
        
     
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



$(document).ready(function() {  
    $('#att_table').DataTable({
        // "serverSide":true,
        // "JQueryUI": true,
       "processing": true,
        "ajax": {
            "url": "/attendence/gettoday", //要抓哪個地方的資料
            "type": "GET", //使用什麼方式抓
            "dataType": 'json', //回傳資料的類型
            "dataSrc": "data",
           // "contentType": "application/json; charset=utf-8",
            // "success": function(data){
            //     console.log("good!")
            //     console.log(data)
            // }, //成功取得回傳時的事件
            // "error": function(){
            //     console.log("資料取得失敗 回去檢討檢討")
            // } //失敗事件
        },
        "columns": [
            { 
                "data": "DATE_S",
                // "render": function (data, type, row, meta) {
                //     return data.substr(0,4)+'/'+data.substr(4,2)+'/'+data.substr(6,2);
                // }
            },
            { "data": "NAME" }, //第一欄使用data中的
            { "data": "EMPID" }, //第二欄使用data中 
            { "data": "ARR" }, //第二欄使用data中
            { "data": "REMARK" ,
          "render": function (data, type, row, meta) {
                    return `<div class="select-editable">
                    <select onchange="this.nextElementSibling.value=this.value" >
                        <option value=""></option>
                        <option value="假（一小）">假（一小）</option>
                        <option value="假（兩小）">假（兩小）</option>
                        <option value="假（兩小）">假（三小）</option>
                        <option value="假（半天）">假（半天）</option>
                        <option value="假（全天）">假（全天）</option>
                    </select>
                    <input type="text" name="format" class="remark" value="`+data+`" />
                </div>`;
                }
        
        }, //第二欄使用data中
            ]//第三欄使用data中的startdate
        
           
    });  
   
}

);  



$(document).on('change', '.select-editable', function(event) {
    // Does some stuff and logs the event to the console
    event.preventDefault();
    var remark = $(this).value;
    var row = $(this).parents('tr')[0];
    var date_s=$($(this).parents('tr')[0].children[0]).text();
    var name=$($(this).parents('tr')[0].children[1]).text();
    var empid=$($(this).parents('tr')[0].children[2]).text();
    
    var remark = $(row.children[4]).find("input").val();
    //console.log(event,row,remark,date_s);
    console.log(date_s,empid,remark);
    edit(date_s,empid,remark);
  });