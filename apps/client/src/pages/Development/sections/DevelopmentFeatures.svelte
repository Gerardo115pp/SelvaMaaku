<script>
    import WixaricaIcon from "@components/icons/wixarica_icon.svelte";
    import { layout_images } from "@stores/layout";
</script>

<section id="smk-development-page-features-section">
    <header id="smk-development-page-features-header" class="design-content-width">
        <div class="wixarica-icons">
            <WixaricaIcon with_accent/>
        </div>
        <hgroup id="smk-dp-header-headliners">
            <h2 id="smk-development-page-features-headline">
                A low density Urban Residential Environment
            </h2>
            <p id="smk-development-page-features-subheadline">
                The architectural project contemplates only 14 houses, that along with the already existing 12 provide generous spaces for ideal gardens to cultivate green areas and reforest this space with Tommy Atkins and Keiit Mango varieties, besides Papelillo Rojo, Prickly Pear, Neem, Eardrop Tree, amongst other native species in the area.
            </p>
        </hgroup>
    </header>
    <div id="smk-development-features-wrapper" class="design-content-width">
        <ul id="smk-dp-features-container" class:adebug={false}>
            <li class="smk-dp-feature-item" id="smk-dp-feature-item-1">
                <div class="wixarica-icons">
                    <WixaricaIcon is_flat_color opacity={0.5}/>
                </div>
                <hgroup>
                    <h3 class="smk-fi-feature-name">Feature name</h3>
                    <p class="smk-fi-feature-description">
                        This is a feature description that needs to be filled with information.
                    </p>
                </hgroup>
            </li>
            <li class="smk-dp-feature-item" id="smk-dp-feature-item-2">
                <div class="wixarica-icons">
                    <WixaricaIcon is_flat_color opacity={0.5}/>
                </div>
                <hgroup>
                    <h3 class="smk-fi-feature-name">Feature name</h3>
                    <p class="smk-fi-feature-description">
                        This is a feature description that needs to be filled with information.
                    </p>
                </hgroup>
            </li>
            <li class="smk-dp-feature-item" id="smk-dp-feature-item-3">
                <div class="wixarica-icons">
                    <WixaricaIcon is_flat_color opacity={0.5}/>
                </div>
                <hgroup>
                    <h3 class="smk-fi-feature-name">Feature name</h3>
                    <p class="smk-fi-feature-description">
                        This is a feature description that needs to be filled with information.
                    </p>
                </hgroup>
            </li>
            <li class="smk-dp-feature-item" id="smk-dp-feature-item-4">
                <div class="wixarica-icons">
                    <WixaricaIcon is_flat_color opacity={0.5}/>
                </div>
                <hgroup>
                    <h3 class="smk-fi-feature-name">Feature name</h3>
                    <p class="smk-fi-feature-description">
                        This is a feature description that needs to be filled with information.
                    </p>
                </hgroup>
            </li>
            <li class="smk-dp-features-image-wrapper"><img src="{layout_images.HOUSE_MODEL_ONE.getUrl(0.3)}" alt="selvamaaku house model"></li>
        </ul>
    </div>
</section>

<style>
    #smk-development-page-features-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: var(--spacing-5) 0;
        row-gap: var(--spacing-5);
    }

    
    /*=============================================
    =            headlines            =
    =============================================*/
    
        header#smk-development-page-features-header {
            width: min(812px, 66.77635%);
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-3);
        }

        hgroup#smk-dp-header-headliners {
            width: 100%;
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-2);
            text-align: center;
        }

        h2#smk-development-page-features-headline {
            font-size: var(--font-size-h3);
            line-height: var(--line-height-9);
        }
    
    /*=====  End of headlines  ======*/

    
    /*=============================================
    =            features            =
    =============================================*/
    
        ul#smk-dp-features-container {
            width: 100%;
            height: 640px;
            display: grid;
            grid-template: 50% 50% / repeat(3, 1fr);
        }

        li.smk-dp-feature-item {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-2);
            padding: var(--spacing-3);

            & hgroup {
                display: flex;
                text-align: center;
                flex-direction: column;
                gap: var(--spacing-1);
            }

            & hgroup h3 {
                font-size: var(--font-size-4);
                line-height: var(--line-height-8);
            }

            & hgroup p {
                font-size: var(--font-size-1);
            }
        }

        #smk-dp-feature-item-1 {
            grid-area: 1 / 1 / span 1 / span 1;
            justify-content: flex-end;
        }

        #smk-dp-feature-item-2 {
            grid-area: 1 / 3 / span 1 / span 1;
            justify-content: flex-end;
        }

        #smk-dp-feature-item-3 {
            grid-area: 2 / 1 / span 1 / span 1;
        }

        #smk-dp-feature-item-4 {
            grid-area: 2 / 3 / span 1 / span 1;
        }

        li.smk-dp-features-image-wrapper {
            grid-area: 1 / 2 / span 2 / span 1;

            & img {
                width: 100%;
                height: 100%;
                object-fit: cover;
                object-position: center;
            }
        }
    
    /*=====  End of features  ======*/
    
    
    
    

</style>