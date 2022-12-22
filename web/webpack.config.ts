/* eslint-disable @typescript-eslint/no-var-requires */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */

import { Configuration, DefinePlugin } from 'webpack';
// import { Configuration, DefinePlugin } from 'webpack';
import { CustomWebpackBrowserSchema, TargetOptions } from '@angular-builders/custom-webpack';
import path from 'path';


/**
 * This is where you define a function that modifies your webpack config
 */
export default (cfg: Configuration, opts: CustomWebpackBrowserSchema, targetOptions: TargetOptions) => {

  const packageName = require('./package.json').name;
  const logo = require("fs").readFileSync("./.profile", 'utf8').replace('$projectname', packageName);
  console.warn(`\n${logo}\n`);
  const entry = cfg.entry! as Record<string, string>;

  // entry['editor.worker'] = 'monaco-editor/esm/vs/editor/editor.worker.js';

  cfg.module?.rules?.push({
    test: /\.(jpe?g|png|svg|gif)/i,
    type: 'asset',
    generator: {
      filename: 'img/[hash][ext][query]' // 局部指定输出位置
    },
    parser: {
      dataUrlCondition: {
        maxSize: 8 * 1024 // 限制于 8kb
      }
    }
  },
    {
      test: /\.txt/,
      type: 'asset/source'
    });

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  // (<any>cfg.resolve!.alias!)['@'] = path.resolve(__dirname, '../src');



  // cfg.module?.rules?.push({
  //   // test: /\.(txt|md)/,
  //   test: (file) => {
  //     if (file.endsWith('.txt')) {
  //       console.warn(file);
  //       return true;
  //     }
  //     return false;
  //   },
  //   type: 'javascript/auto',
  //   use: [{
  //     loader: require.resolve('raw-loader'),
  //     options: {
  //       esModule: false
  //     }
  //   }],
  //   // loader: ,
  //   // resolve: {
  //   //   fullySpecified: true
  //   // },
  //   enforce: "pre"
  // });

  // console.warn(JSON.stringify(cfg));

  // cfg.plugins!.push(
  //   new HtmlWebpackPlugin({
  //     filename: 'footer.html',
  //     template: 'src/footer-template.html',
  //   }),
  //   new DefinePlugin({
  //     APP_VERSION: JSON.stringify(version),
  //   })
  // );

  return cfg;
};



// https://webpack.js.org/guides/asset-modules/