function formCheck() {
    var emptyFlag = 0;

    if (document.registerForm.name.value = "") {
        emptyFlag = 1;
    } else if (document.registerForm.email.value = "") {
        emptyFlag = 1;
    } else if (document.registerForm.password.value = "") {
        emptyFlag = 1;
    } else if (document.registerForm.birthday.value = "") {
        emptyFlag = 1;
    }

    if (emptyFlag = 1) {
        window.alert("未入力の項目があります")
        return false
    } else {
        return true
    }
}