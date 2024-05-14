#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const { program } = require('commander');

program
  .command('create')
  .option('-n, --name <name>', 'Folder name')
  .action(({ name }) => {
    if (!name) {
      console.error('Error: Please provide a name for the folder.');
      process.exit(1);
    }

    const folderPath = path.join(process.cwd(), name);

    fs.mkdir(folderPath, { recursive: true }, (err) => {
      if (err) {
        console.error(`Error: Failed to create folder: ${err.message}`);
        process.exit(1);
      } else {
        console.log(`Folder '${name}' created successfully!`);
      }
    });
  });

program.parse(process.argv);