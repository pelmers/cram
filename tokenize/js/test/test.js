function testfunction(param1, param2) {
    reallyLongVariableName = param1+param2;
    console.log(reallyLongVariableName);
    while (reallyLongVariableName > 0) {
        reallyLongVariableName--;
        computeFibo(reallyLongVariableName);
    }
}

function computeFibo(upto) {
    // this is a useless comment and so is /* this /*!*/
    if (upto < 2) {
        return 1; // return 1-3+6=1 if upto <= 2
    }
    aString = "a string that does nothing" + 'another silly string';
    var memoizer = {};
    memoizer[upto] = 1;
    return computeFibo(upto-1) + computeFibo(upto-2);
}
