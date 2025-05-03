import React from 'react'

type Props = {}

function NotFound({}: Props) {
  return (
    <div className='flex flex-col h-full items-center justify-center'>
        <h1 className='text-2xl'>Not Found</h1>
    </div>
  )
}

export default NotFound