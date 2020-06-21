import { machine, DiaryMachine } from './machine';
import { interpret } from 'xstate';

export function interpretMachine(): DiaryMachine {
  return interpret(machine.withConfig({})) as DiaryMachine;
}
