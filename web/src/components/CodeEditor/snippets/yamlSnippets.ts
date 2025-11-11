export const yamlSnippets = (monaco) => {
    return [
        {
            label: 'true',
            kind: monaco.languages.CompletionItemKind.Keyword,
            insertText: `true\${1:}`,
            insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
            documentation: 'true'
        },
        {
            label: 'false',
            kind: monaco.languages.CompletionItemKind.Keyword,
            insertText: `false\${1:}`,
            insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
            documentation: 'false'
        }
    ];
}