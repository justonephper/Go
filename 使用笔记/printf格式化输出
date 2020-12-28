基本介绍：

    %v	the value in a default format
	    when printing structs, the plus flag (%+v) adds field names
    %#v	a Go-syntax representation of the value
    %T	a Go-syntax representation of the type of the value
    %%	a literal percent sign; consumes no value
    
 The default format for %v is:
    bool:                    %t
    int, int8 etc.:          %d
    uint, uint8 etc.:        %d, %#x if printed with %#v
    float32, complex64, etc: %g
    string:                  %s
    chan:                    %p
    pointer:                 %p

 Boolean:
    %t	the word true or false
    
 Integer:
    %b	decimalless scientific notation with exponent a power of two,
	    in the manner of strconv.FormatFloat with the 'b' format,
	    e.g. -123456p-78
    %e	scientific notation, e.g. -1.234456e+78
    %E	scientific notation, e.g. -1.234456E+78
    %f	decimal point but no exponent, e.g. 123.456
    %F	synonym for %f
    %g	%e for large exponents, %f otherwise
    %G	%E for large exponents, %F otherwise
    
 String and slice of bytes (treated equivalently with these verbs):
    %s	the uninterpreted bytes of the string or slice
    %q	a double-quoted string safely escaped with Go syntax
    %x	base 16, lower-case, two characters per byte
    %X	base 16, upper-case, two characters per byte
    
 Pointer:
    %p	base 16 notation, with leading 0x
    
    
    
    
