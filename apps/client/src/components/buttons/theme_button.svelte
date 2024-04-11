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

        /**
         * Whether or not to use target="_blank" in the <a> tag
         * @type {boolean}
         * @default false
         */
        export let open_in_new_tab = false;
    
    /*=====  End of Properties  ======*/
    
    
</script>

{#if href != null}
    <a href="{href}"
        target={open_in_new_tab ? "_blank" : "_self"}
        rel={open_in_new_tab ? "noopener noreferrer" : ""}
    >
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
