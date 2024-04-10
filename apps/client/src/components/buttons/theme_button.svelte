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

        /**
         * Sets the button class to be button-3 which is a variation of button-2
         * @type {boolean}
         * @default false
         */
        export let is_button_three = false;

        /**
         * if set, the button will be wrapped around and <a> tag and have a navigation role
         * @type {string}
         */
        export let href;
    
    /*=====  End of Properties  ======*/
    
    
</script>

{#if href != null}
    <a href="{href}">
        <button 
            type="button"
            class:button-2={!is_button_one && !is_button_three}
            class:button-1={is_button_one}
            class:button-3={is_button_three}
            role="navigation"
        >
            <span>  
                {text}
            </span>
            <Arrow />
        </button>
    </a>
{:else}
    <button 
        type="button"
        class:button-2={!is_button_one && !is_button_three}
        class:button-1={is_button_one}
        class:button-3={is_button_three}
    >
        <span>  
            {text}
        </span>
        <Arrow />
    </button>
{/if}
