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

        /**
         * Whether or not to use a curtain animation on active(click)
         * @type {boolean}
         * @default true
         */
        export let animate_active_curtain = true;
    
    /*=====  End of Properties  ======*/
    
    
</script>

{#if href != null}
    <a href="{href}"
        target={open_in_new_tab ? "_blank" : "_self"}
        rel={open_in_new_tab ? "noopener noreferrer" : ""}
    >
        <button 
            class="smk-theme-button"
            type="button"
            class:button-2={!is_button_one && !is_button_three}
            class:button-1={is_button_one}
            class:button-3={is_button_three}
            class:button-active-curtain={animate_active_curtain}
            role="navigation"
        >
            <span class="smk-theme-button-curtain"></span>
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

<style>
    .smk-theme-button {
        overflow: hidden;
        position: relative;
        isolation: isolate;
    }

    
    /*=============================================
    =            Curtain animation            =
    =============================================*/

        @keyframes slide-curtain {
            50% {
                scale: 1 1;
                opacity: 1;
            }
            100% {
                scale: 1 1;
                opacity: 0.7;
            }
        }

        span.smk-theme-button-curtain {
            display: none;
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            scale: 0 1;
            transform-origin: left;
            background: var(--theme-hero-background);
            z-index: var(--z-index-b-1);
            transition: opacity 0.3s ease;
        }

        button.button-active-curtain.button-1 span.smk-theme-button-curtain {
            background: var(--accent-orange-active);
        }

        button.button-active-curtain span.smk-theme-button-curtain {
            display: block;
        }

        button.button-active-curtain:hover span.smk-theme-button-curtain {
            animation: slide-curtain 0.6s both;
        }        
    
    /*=====  End of Curtain animation  ======*/
    
    


</style>
