Cuz all the cool kidz cram.

```
go get github.com/pelmers/cram
```

Process:
========
- [main.go](https://github.com/pelmers/cram/blob/master/main.go) parses flags and sets options.
- A [tokenizer](https://github.com/pelmers/cram/tree/master/tokenize) splits input into whitespace-dempotent entities (i.e. amount of whitespace between consecutive tokens will not affect semantics).
- Optionally, a renamer will parse the tokens and minify variable names.
- Then a [reshaper](https://github.com/pelmers/cram/tree/master/shapes) is dispatched to put the tokens back together again into fancy shapes.

Example:
=======
                         function  
                        wN(  Xo  ){ 
                      "use strict"  ;
                    function io( nv){Xo
                   .JY(nv, 'font-weight'
                 ,'bold') ;   Xo . JY( nv,
                'font-weight'  );Xo .JY(nv, 
              'stroke-width','0px') ; if (Xo.
             mm ()  .property(nv,'type'  )  ==
            'ACTIVITY_SPAWN'){Xo.Rc(nv,  'fill',
          'rgb(161,217,155)');}}function Xg(nv,UW 
         ){ for (var pm=0;pm<UW .length;PN++ ){var 
       Se=UW[pm] ; if (Xo.mm().edge_property( nv ,Se,
      'type') === 'join'   ){Xo.lM(nv ,Se , 'cursor' ,
    'pointer'  ); }    }  } Xo  .  Pk(    io,  Xg) ;   }

