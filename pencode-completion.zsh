compdef _pencode pencode

function _pencode {
    _arguments "*: :($(pencode -list))"
}