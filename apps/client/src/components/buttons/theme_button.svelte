<script>
    import Arrow from "@components/icons/arrow.svelte";

    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        /**
         * Text to display in the button
         * @type {string}
         */
        export let text; 

        /**
         * Controls the class of the button. if false, the button class will be button-2, if true, the button class will be button-1
         * @type {boolean}
         * @default false
         */
        export let is_button_one = false;
    
    /*=====  End of Properties  ======*/
    
    
</script>

<button 
    class:button-2={!is_button_one}
    class:button-1={is_button_one}
>
    <span>  
        {text}
    </span>
    <Arrow />
</button>