import type { assert } from "console";

interface User {
    id: number;
    username: string;
    email: string;
    password: string;
    yob: number;
    createdAt: string;
    updatedAt: string;
}

interface CodeFile {
    name: string;
    content: string;
    readonly: boolean;
}

interface Lesson {
    id: number;
    name: string;
    slug: string;
    key_name: string;
    description: string;
    files: CodeFile[];
    dialogs: Dialog[];
    helpers: string[];
    asserts: string[];
    next_lesson: string;
    position: number;
    completed: boolean;
    createdAt: string;
    updatedAt: string;
}

interface Course {
    id: number;
    name: string;
    slug: string;
    key_name: string;
    description: string;
    lessons: Lesson[];
    createdAt: string;
    updatedAt: string;
}

interface Lang {
    id: number;
    name: string;
    slug: string;
    key_name: string;
    description: string;
    courses: Course[];
    createdAt: string;
    updatedAt: string;
}

interface Dialog {
    key_name: string;
    trigger: string;
}

interface Alert {
    id: string;
    type: string;
    message: string;
}

type AlertType = 'success' | 'error' | 'warning' | 'info';

export type { User, CodeFile, Lesson, Dialog, Alert, AlertType, Course, Lang };
