// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const StudentRole = {
    Student: "Student",
    Admin: "Admin",
} as const;

/**
 * Student Role
 */
export type StudentRole = (typeof StudentRole)[keyof typeof StudentRole];
