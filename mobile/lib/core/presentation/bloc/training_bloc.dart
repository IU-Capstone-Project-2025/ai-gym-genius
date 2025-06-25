import 'package:flutter_bloc/flutter_bloc.dart';

import '/core/domain/entities/exercise_entity.dart';
import '/core/domain/repositories/training_repository.dart';
import 'training_event.dart';
import 'training_state.dart';

class TrainingBloc extends Bloc<TrainingEvent, TrainingState> {
  final TrainingRepository workoutRepository;
  TrainingBloc({required this.workoutRepository}) : super(TrainingState()) {
    on<AddExercise>(_onAddExercise);
    on<RemoveExercise>(_onRemoveExercise);
    on<SubmitTraining>(_onSubmitTraining);
  }

  void _onAddExercise(AddExercise event, Emitter<TrainingState> emit) {
    // Make state initial first (clean dirty)
    emit(state.copyWith(addStatus: AddExerciseStatus.initial));

    final exists = state.exercises.any(
      (element) => element.exerciseInfo.id == event.info.id,
    );
    if (exists) {
      emit(state.copyWith(addStatus: AddExerciseStatus.duplicate));
      return;
    }
    final exercise = ExerciseEntity(exerciseInfo: event.info, sets: []);
    emit(state.copyWith(
        exercises: state.exercises + [exercise],
        addStatus: AddExerciseStatus.success));
  }

  void _onRemoveExercise(RemoveExercise event, Emitter<TrainingState> emit) {
    emit(
      state.copyWith(
        exercises: state.exercises
            .where((obj) => obj.exerciseInfo.id != event.exerciseID)
            .toList(),
      ),
    );
  }

  void _onSubmitTraining(
      SubmitTraining event, Emitter<TrainingState> emit) async {}
}
