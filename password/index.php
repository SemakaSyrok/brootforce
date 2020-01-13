<?php

if( isset($_POST['password']) && $_POST['login'] ) {
    
    if($_POST['password'] == 'sefrvewve' && $_POST['login'] == 'login') {
        echo 'You are logged in!';
        die();
    }
    echo $_POST['password'] . '  ' . $_POST['login'];   
}

?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Пароль</title>
</head>
<body>


<form action="" method="POST">

    <div class="">
        <label for="login">Login</label>
        <input type="text" id="login" name="login" >
    </div>

    <br>

    <div class="">
        <label for="password">Password</label>
        <input type="text" id="password" name="password" >
    </div>

    <br>

    <button name="send" >Send</button>

</form>

</body>
</html>