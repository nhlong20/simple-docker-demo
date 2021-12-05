import React from 'react';

function Course({ course, handleClick }) {
    return (
        <div className="p-6 bg-white rounded-xl shadow-md mt-6">
            <a className="text-xl text-blue-900 font-semibold hover:text-red-700 cursor-pointer">
                {course.code} - {course.title}
            </a>
            <div className="py-1">
                {course.department} - {course.credit} tín chỉ
            </div>
            <div>
                {course.description}
            </div>
        </div>
    )
}

export default Course