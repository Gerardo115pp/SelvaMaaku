<script>
    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        export let video_url = '';

        /**
         * The resource that will be used as the poster for the video.
         * @type {import('@models/ImageResources').ImageResource}
         * @default null
         */
        export let poster_media;

        /**
         * The percentage of the viewport that the video poster will be loaded for.
         * @type {number}
         * @default 0.5
        */
        export let poster_threshold = 0.5;

        //* Style properties
        
        /**
         * The width of the video player.
         * @type {string}
         * @default "max-content"
         */
        export let player_width = "max-content";
        /**
         * The height of the video player.
         * @type {string}
         * @default "auto"
         */
        export let player_height = "auto";
    
    /*=====  End of Properties  ======*/
    
    
</script>

<div 
    id="video-player-component"
    style:width={player_width}
    style:height={player_height}
>
    {#if video_url === ''}
        <div id="video-player-poster-wrapper">
            <!-- TODO: Change the video alt for something descriptive and SEO helpful -->
            <img src="{poster_media.getUrl(poster_threshold)}" alt="This video is amazing">
        </div>
    {/if}
</div>

<style>
    #video-player-component {
        container-type: size;
    }

    #video-player-poster-wrapper img {
        width: 100cqw;
        height: 100cqh;
        object-fit: cover;
        object-position: center;
    }
</style>