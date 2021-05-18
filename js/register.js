function formCheck() {
    var emptyFlag = 0;

    if (registerForm.name.value == "") {
        emptyFlag = 1;
    } else if (registerForm.email.value == "") {
        emptyFlag = 1;
    } else if (registerForm.password.value == "") {
        emptyFlag = 1;
    } else if (registerForm.birthday.value == "") {
        emptyFlag = 1;
    }

    if (emptyFlag == 1) {
        window.alert("必要事項が入力されていません")
        return false;
    } else {
        return true;
    }
}