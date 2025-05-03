import { SearchIcon, Type } from 'lucide-react'
import React from 'react'

type Props = {}

const Search = (props: Props) => {
  return (
    <div className='flex items-center border-1 px-3 border-neutral-300 rounded-xl'>
        <SearchIcon  size={19} color='black' />
        <input className='outline-none border-none p-2' type='text' placeholder='search ...' />
        <Type className='p-2 rounded-md cursor-default bg-neutral-100 outline-1 duration-300 outline-neutral-300 hover:bg-neutral-300' size={29} color='black' />
    </div>
  )
}

export default Search