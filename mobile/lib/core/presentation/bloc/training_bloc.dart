import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:gym_genius/core/domain/entities/workout_entity.dart';

import '/core/domain/entities/exercise_entity.dart';
import '../../domain/repositories/workout_repository.dart';
import 'training_event.dart';
import 'training_state.dart';

class TrainingBloc extends Bloc<TrainingEvent, TrainingState> {
  final WorkoutRepository workoutRepository;
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
      SubmitTraining event, Emitter<TrainingState> emit) async {
    emit(state.copyWith(submitStatus: SubmitTrainingStatus.initial));

    // First we check if some sets are empty
    final existsEmpty = state.exercises.any(
      (exercise) => exercise.sets.isEmpty,
    );
    if (existsEmpty) {
      emit(state.copyWith(submitStatus: SubmitTrainingStatus.failureEmptySets));
      return;
    }

    final workout = WorkoutEntity(
      id: -1, // Do not need an ID really
      duration: event.workoutDuration,
      startTime: DateTime.now().subtract(event.workoutDuration),
      exercises: state.exercises,
    );

    emit(state.copyWith(submitStatus: SubmitTrainingStatus.loading));
    try {
      workoutRepository.saveWorkout(workout);
    } catch (e) {
      emit(state.copyWith(submitStatus: SubmitTrainingStatus.failureWritingDB));
    }

    // FINAL STATE
    emit(state.copyWith(
      submitStatus: SubmitTrainingStatus.success,
      workout: workout,
    ));
  }
}
