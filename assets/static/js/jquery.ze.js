
jQuery.fn.extend({
    template : function(data){
        var tmpl_data = $(this).html();

        var fn = new Function("obj",
                     "var p=[];" +

                     // Introduce the data as local variables using with(){}
                     "with(obj){p.push('" +

                     // Convert the template into pure JavaScript
                    tmpl_data 
                     .replace(/[\r\t\n]/g, " ")
                     .split("<%").join("\t")
                     .replace(/(^|%>)[^\t]*?(\t|$)/g, function(){return arguments[0].split("'").join("\\'");})
                     .replace(/\t==(.*?)%>/g,"',$1,'")
                     .replace(/\t=(.*?)%>/g, "',(($1)+'').replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/\"/g,'&quot;'),'")
                     .split("\t").join("');")
                     .split("%>").join("p.push('")
                     + "');}return p.join('');");
        return fn( data );
            
    }

});

/**
 * jQuery.bottom
 * Dual licensed under MIT and GPL.
 * Date: 2010-04-25
 *
 * @description Trigger the bottom event when the user has scrolled to the bottom of an element
 * @author Jim Yi
 * @version 1.0
 *
 * @id jQuery.fn.bottom
 * @param {Object} settings Hash of settings.
 * @return {jQuery} Returns the same jQuery object for chaining.
 *
 */



(function($){
    $.fn.bottom = function(options) {

        var defaults = {
            proximity: 0
        };

        var options = $.extend(defaults, options);

        return this.each(function() {
            var obj = this;
            $(obj).bind("scroll", function() {
                if (obj == window) {
                    scrollHeight = $(document).height();
                }
                else {
                    scrollHeight = $(obj)[0].scrollHeight;
                }
            scrollPosition = $(obj).height() + $(obj).scrollTop();
            if ( (scrollHeight - scrollPosition) / scrollHeight <= options.proximity) {
                $(obj).trigger("bottom");
            }
            });

            return false;
        });
    };
})(jQuery);


// firebug console..
if (window.console && window.console.log ) {
    window.log = window.console.log
} else {
    window.console = {
        log: function () {}
    }
}

