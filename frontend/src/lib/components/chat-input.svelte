<script lang="ts">
    import IconGIF from "~icons/fluent/gif-16-regular"
    import IconSmiley from "~icons/fluent/emoji-16-regular"
</script>

<div class="flex items-center h-24 mx-3 border-t border-gray-600 bg-gray-750">
    <div class="flex-1">
        <input
            type="text"
            class="w-full h-10 px-3 py-2 text-sm text-gray-200 rounded-tl rounded-bl bg-slate-800 focus:outline-none"
        />
    </div>
    <div class="flex items-center h-10 px-2 py-2 rounded-r bg-slate-800">
        <button class="h-10 ml-3 text-gray-500 rounded-l bg-slate-800 hover:text-white">
            <IconGIF />
        </button>
        <button class="h-10 ml-3 text-gray-500 rounded-l bg-slate-800 hover:text-white">
            <IconSmiley />
        </button>
    </div>
</div>
