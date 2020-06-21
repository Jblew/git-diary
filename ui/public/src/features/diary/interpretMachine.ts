import { machine, DiaryMachine } from './machine';
import { interpret } from 'xstate';
import { readDiary, publishEntry } from './services';

export function interpretMachine(): DiaryMachine {
  return interpret(machine.withConfig({
    services: {
      readDiary,
      publishEntry: (ctx) => publishEntry(ctx.publishText),
    },
  })) as DiaryMachine;
}
