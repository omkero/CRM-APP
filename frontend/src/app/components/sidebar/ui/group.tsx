import React from 'react'

type Props = {
    GroupTitle: string
    children: React.ReactNode
}

function SideBarGroup({GroupTitle, children}: Props) {
  return (
    <div className='flex w-full flex-col gap-4'>
        <div className='px-3'>
        <h1 className='text-gray-500'>{GroupTitle}</h1>
        </div>
        <div className='flex flex-col gap-2 items-center w-full'>
            {children}
        </div>
    </div>
  )
}

export default SideBarGroup