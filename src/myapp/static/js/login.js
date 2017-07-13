$(function(){
	$('#submit').click(function(){
		//>>获取用户名和密码
		var user = $('input[name = username]').val()

		var password = $('input[name = password]').val()
	
		if(user ==""){
			alert("用户名不能为空")
			return false
		}

		if(password == ""){
			alert("密码不能为空")
			return false
		}
		
		$.ajax({
			'type':'post',
			'dataType':'json',
			'url':location.protocol+'//'+window.location.host+'/check',
			'data':{'username':user,'password':password},
			success:function(e){
				console.log(e)
			}
		})
	})
})